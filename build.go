// Copyright (C) 2014 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build ignore
// +build ignore

package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"time"

	buildpkg "github.com/expgo/net/build"
)

var (
	goarch         string
	goos           string
	noupgrade      bool
	version        string
	goCmd          string
	race           bool
	debug          = os.Getenv("BUILDDEBUG") != ""
	extraTags      string
	installSuffix  string
	pkgdir         string
	cc             string
	run            string
	benchRun       string
	buildOut       string
	debugBinary    bool
	coverage       bool
	long           bool
	timeout        = "120s"
	longTimeout    = "600s"
	numVersions    = 5
	withNextGenGUI = os.Getenv("BUILD_NEXT_GEN_GUI") != ""
)

type target struct {
	name              string
	debname           string
	debdeps           []string
	debpre            string
	description       string
	buildPkgs         []string
	binaryName        string
	archiveFiles      []archiveFile
	systemdService    string
	installationFiles []archiveFile
	tags              []string
}

type archiveFile struct {
	src  string
	dst  string
	perm os.FileMode
}

func main() {
	log.SetFlags(0)

	parseFlags()

	if debug {
		t0 := time.Now()
		defer func() {
			log.Println("... build completed in", time.Since(t0))
		}()
	}

	// Invoking build.go with no parameters at all builds everything (incrementally),
	// which is what you want for maximum error checking during development.
	if flag.NArg() == 0 {
		runCommand("proto")
	} else {
		runCommand(flag.Arg(0))
	}
}

func runCommand(cmd string) {
	var tags []string
	if noupgrade {
		tags = []string{"noupgrade"}
	}
	tags = append(tags, strings.Fields(extraTags)...)

	switch cmd {
	case "test":
		test(strings.Fields(extraTags), "github.com/syncthing/syncthing/lib/...", "github.com/syncthing/syncthing/cmd/...")

	case "bench":
		bench(strings.Fields(extraTags), "github.com/syncthing/syncthing/lib/...", "github.com/syncthing/syncthing/cmd/...")

	case "integration":
		integration(false)

	case "integrationbench":
		integration(true)

	case "assets":
		rebuildAssets()

	case "update-deps":
		updateDependencies()

	case "proto":
		proto()

	case "translate":
		translate()

	case "transifex":
		transifex()

	case "vet":
		metalintShort()

	case "lint":
		metalintShort()

	case "metalint":
		metalint()

	case "version":
		fmt.Println(getVersion())

	case "changelog":
		vers, err := currentAndLatestVersions(numVersions)
		if err != nil {
			log.Fatal(err)
		}
		for _, ver := range vers {
			underline := strings.Repeat("=", len(ver))
			msg, err := tagMessage(ver)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n%s\n\n%s\n\n", ver, underline, msg)
		}

	default:
		log.Fatalf("Unknown command %q", cmd)
	}
}

func parseFlags() {
	flag.StringVar(&goarch, "goarch", runtime.GOARCH, "GOARCH")
	flag.StringVar(&goos, "goos", runtime.GOOS, "GOOS")
	flag.StringVar(&goCmd, "gocmd", "go", "Specify `go` command")
	flag.BoolVar(&noupgrade, "no-upgrade", noupgrade, "Disable upgrade functionality")
	flag.StringVar(&version, "version", getVersion(), "Set compiled in version string")
	flag.BoolVar(&race, "race", race, "Use race detector")
	flag.StringVar(&extraTags, "tags", extraTags, "Extra tags, space separated")
	flag.StringVar(&installSuffix, "installsuffix", installSuffix, "Install suffix, optional")
	flag.StringVar(&pkgdir, "pkgdir", "", "Set -pkgdir parameter for `go build`")
	flag.StringVar(&cc, "cc", os.Getenv("CC"), "Set CC environment variable for `go build`")
	flag.BoolVar(&debugBinary, "debug-binary", debugBinary, "Create unoptimized binary to use with delve, set -gcflags='-N -l' and omit -ldflags")
	flag.BoolVar(&coverage, "coverage", coverage, "Write coverage profile of tests to coverage.txt")
	flag.BoolVar(&long, "long", long, "Run tests without the -short flag")
	flag.IntVar(&numVersions, "num-versions", numVersions, "Number of versions for changelog command")
	flag.StringVar(&run, "run", "", "Specify which tests to run")
	flag.StringVar(&benchRun, "bench", "", "Specify which benchmarks to run")
	flag.BoolVar(&withNextGenGUI, "with-next-gen-gui", withNextGenGUI, "Also build 'newgui'")
	flag.StringVar(&buildOut, "build-out", "", "Set the '-o' value for 'go build'")
	flag.Parse()
}

func test(tags []string, pkgs ...string) {
	lazyRebuildAssets()

	tags = append(tags, "purego")
	args := []string{"test", "-tags", strings.Join(tags, " ")}
	if long {
		timeout = longTimeout
	} else {
		args = append(args, "-short")
	}
	args = append(args, "-timeout", timeout)

	if runtime.GOARCH == "amd64" {
		switch runtime.GOOS {
		case buildpkg.Darwin, buildpkg.Linux, buildpkg.FreeBSD: // , "windows": # See https://github.com/golang/go/issues/27089
			args = append(args, "-race")
		}
	}

	if coverage {
		args = append(args, "-covermode", "atomic", "-coverprofile", "coverage.txt", "-coverpkg", strings.Join(pkgs, ","))
	}

	args = append(args, runArgs()...)

	runPrint(goCmd, append(args, pkgs...)...)
}

func bench(tags []string, pkgs ...string) {
	lazyRebuildAssets()
	args := append([]string{"test", "-run", "NONE", "-tags", strings.Join(tags, " ")}, benchArgs()...)
	runPrint(goCmd, append(args, pkgs...)...)
}

func integration(bench bool) {
	lazyRebuildAssets()
	args := []string{"test", "-v", "-timeout", "60m", "-tags"}
	tags := "purego,integration"
	if bench {
		tags += ",benchmark"
	}
	args = append(args, tags)
	args = append(args, runArgs()...)
	if bench {
		if run == "" {
			args = append(args, "-run", "Benchmark")
		}
		args = append(args, benchArgs()...)
	}
	args = append(args, "./test")
	fmt.Println(args)
	runPrint(goCmd, args...)
}

func runArgs() []string {
	if run == "" {
		return nil
	}
	return []string{"-run", run}
}

func benchArgs() []string {
	if benchRun == "" {
		return []string{"-bench", "."}
	}
	return []string{"-bench", benchRun}
}

func setBuildEnvVars() {
	os.Setenv("GOOS", goos)
	os.Setenv("GOARCH", goarch)
	os.Setenv("CC", cc)
	if os.Getenv("CGO_ENABLED") == "" {
		switch goos {
		case "darwin", "solaris":
		default:
			os.Setenv("CGO_ENABLED", "0")
		}
	}
}

func appendParameters(args []string, tags []string, pkgs ...string) []string {
	if pkgdir != "" {
		args = append(args, "-pkgdir", pkgdir)
	}
	if len(tags) > 0 {
		args = append(args, "-tags", strings.Join(tags, " "))
	}
	if installSuffix != "" {
		args = append(args, "-installsuffix", installSuffix)
	}
	if race {
		args = append(args, "-race")
	}

	if !debugBinary {
		// Regular binaries get version tagged and skip some debug symbols
		args = append(args, "-trimpath", "-ldflags", ldflags(tags))
	} else {
		// -gcflags to disable optimizations and inlining. Skip -ldflags
		// because `Could not launch program: decoding dwarf section info at
		// offset 0x0: too short` on 'dlv exec ...' see
		// https://github.com/go-delve/delve/issues/79
		args = append(args, "-gcflags", "all=-N -l")
	}

	return append(args, pkgs...)
}

func createPostInstScript(target target) (string, error) {
	scriptname := filepath.Join("script", "deb-post-inst.template")
	t, err := template.ParseFiles(scriptname)
	if err != nil {
		return "", err
	}
	scriptname = strings.TrimSuffix(scriptname, ".template")
	w, err := os.Create(scriptname)
	if err != nil {
		return "", err
	}
	defer w.Close()
	if err = t.Execute(w, struct {
		Service, Command string
	}{
		target.systemdService, target.binaryName,
	}); err != nil {
		return "", err
	}
	return scriptname, nil
}

func shouldBuildSyso(dir string) (string, error) {
	type M map[string]interface{}
	version := getVersion()
	version = strings.TrimPrefix(version, "v")
	major, minor, patch := semanticVersion()
	bs, err := json.Marshal(M{
		"FixedFileInfo": M{
			"FileVersion": M{
				"Major": major,
				"Minor": minor,
				"Patch": patch,
			},
			"ProductVersion": M{
				"Major": major,
				"Minor": minor,
				"Patch": patch,
			},
		},
		"StringFileInfo": M{
			"CompanyName":      "The Syncthing Authors",
			"FileDescription":  "Syncthing - Open Source Continuous File Synchronization",
			"FileVersion":      version,
			"InternalName":     "syncthing",
			"LegalCopyright":   "The Syncthing Authors",
			"OriginalFilename": "syncthing",
			"ProductName":      "Syncthing",
			"ProductVersion":   version,
		},
		"IconPath": "assets/logo.ico",
	})
	if err != nil {
		return "", err
	}

	jsonPath := filepath.Join(dir, "versioninfo.json")
	err = os.WriteFile(jsonPath, bs, 0644)
	if err != nil {
		return "", errors.New("failed to create " + jsonPath + ": " + err.Error())
	}

	defer func() {
		if err := os.Remove(jsonPath); err != nil {
			log.Printf("Warning: unable to remove generated %s: %v. Please remove it manually.", jsonPath, err)
		}
	}()

	sysoPath := filepath.Join(dir, "cmd", "syncthing", "resource.syso")

	// See https://github.com/josephspurrier/goversioninfo#command-line-flags
	armOption := ""
	if strings.Contains(goarch, "arm") {
		armOption = "-arm=true"
	}

	if _, err := runError("goversioninfo", "-o", sysoPath, armOption); err != nil {
		return "", errors.New("failed to create " + sysoPath + ": " + err.Error())
	}

	return sysoPath, nil
}

func shouldCleanupSyso(sysoFilePath string) {
	if sysoFilePath == "" {
		return
	}
	if err := os.Remove(sysoFilePath); err != nil {
		log.Printf("Warning: unable to remove generated %s: %v. Please remove it manually.", sysoFilePath, err)
	}
}

// copyFile copies a file from src to dst, ensuring the containing directory
// exists. The permission bits are copied as well. If dst already exists and
// the contents are identical to src the modification time is not updated.
func copyFile(src, dst string, perm os.FileMode) error {
	in, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	out, err := os.ReadFile(dst)
	if err != nil {
		// The destination probably doesn't exist, we should create
		// it.
		goto copy
	}

	if bytes.Equal(in, out) {
		// The permission bits may have changed without the contents
		// changing so we always mirror them.
		os.Chmod(dst, perm)
		return nil
	}

copy:
	os.MkdirAll(filepath.Dir(dst), 0777)
	if err := os.WriteFile(dst, in, perm); err != nil {
		return err
	}

	return nil
}

func listFiles(dir string) []string {
	var res []string
	filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.Mode().IsRegular() {
			res = append(res, path)
		}
		return nil
	})
	return res
}

func rebuildAssets() {
	os.Setenv("SOURCE_DATE_EPOCH", fmt.Sprint(buildStamp()))
	runPrint(goCmd, "generate", "github.com/syncthing/syncthing/lib/api/auto", "github.com/syncthing/syncthing/cmd/strelaypoolsrv/auto")
}

func lazyRebuildAssets() {
	shouldRebuild := shouldRebuildAssets("lib/api/auto/gui.files.go", "gui") ||
		shouldRebuildAssets("cmd/strelaypoolsrv/auto/gui.files.go", "cmd/strelaypoolsrv/gui")

	if withNextGenGUI {
		shouldRebuild = buildNextGenGUI() || shouldRebuild
	}

	if shouldRebuild {
		rebuildAssets()
	}
}

func buildNextGenGUI() bool {
	// Check if we need to run the npm process, and if so also set the flag
	// to rebuild Go assets afterwards. The index.html is regenerated every
	// time by the build process. This assumes the new GUI ends up in
	// next-gen-gui/dist/next-gen-gui.

	if !shouldRebuildAssets("gui/next-gen-gui/index.html", "next-gen-gui") {
		// The GUI is up to date.
		return false
	}

	runPrintInDir("next-gen-gui", "npm", "install")
	runPrintInDir("next-gen-gui", "npm", "run", "build", "--", "--prod", "--subresource-integrity")

	rmr("gui/tech-ui")

	for _, src := range listFiles("next-gen-gui/dist") {
		rel, _ := filepath.Rel("next-gen-gui/dist", src)
		dst := filepath.Join("gui", rel)
		if err := copyFile(src, dst, 0644); err != nil {
			fmt.Println("copy:", err)
			os.Exit(1)
		}
	}

	return true
}

func shouldRebuildAssets(target, srcdir string) bool {
	info, err := os.Stat(target)
	if err != nil {
		// If the file doesn't exist, we must rebuild it
		return true
	}

	// Check if any of the files in gui/ are newer than the asset file. If
	// so we should rebuild it.
	currentBuild := info.ModTime()
	assetsAreNewer := false
	stop := errors.New("no need to iterate further")
	filepath.Walk(srcdir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.ModTime().After(currentBuild) {
			assetsAreNewer = true
			return stop
		}
		return nil
	})

	return assetsAreNewer
}

func updateDependencies() {
	runPrint(goCmd, "get", "-u", "./cmd/...")
	runPrint(goCmd, "mod", "tidy", "-go=1.17", "-compat=1.17")

	// We might have updated the protobuf package and should regenerate to match.
	proto()
}

func proto() {
	pv := protobufVersion()
	repo := "https://github.com/gogo/protobuf.git"
	path := filepath.Join("repos", "protobuf")

	runPrint(goCmd, "install", fmt.Sprintf("github.com/gogo/protobuf/protoc-gen-gogofast@%v", pv))
	os.MkdirAll("repos", 0755)

	if _, err := os.Stat(path); err != nil {
		runPrint("git", "clone", repo, path)
	} else {
		runPrintInDir(path, "git", "fetch")
	}
	runPrintInDir(path, "git", "checkout", pv)

	runPrint(goCmd, "generate", "proto/generate.go")
}

func translate() {
	os.Chdir("gui/default/assets/lang")
	runPipe("lang-en-new.json", goCmd, "run", "../../../../script/translate.go", "lang-en.json", "../../../")
	os.Remove("lang-en.json")
	err := os.Rename("lang-en-new.json", "lang-en.json")
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir("../../../..")
}

func transifex() {
	os.Chdir("gui/default/assets/lang")
	runPrint(goCmd, "run", "../../../../script/transifexdl.go")
}

func ldflags(tags []string) string {
	b := new(strings.Builder)
	b.WriteString("-w")
	fmt.Fprintf(b, " -X github.com/syncthing/syncthing/lib/build.Version=%s", version)
	fmt.Fprintf(b, " -X github.com/syncthing/syncthing/lib/build.Stamp=%d", buildStamp())
	fmt.Fprintf(b, " -X github.com/syncthing/syncthing/lib/build.User=%s", buildUser())
	fmt.Fprintf(b, " -X github.com/syncthing/syncthing/lib/build.Host=%s", buildHost())
	fmt.Fprintf(b, " -X github.com/syncthing/syncthing/lib/build.Tags=%s", strings.Join(tags, ","))
	if v := os.Getenv("EXTRA_LDFLAGS"); v != "" {
		fmt.Fprintf(b, " %s", v)
	}
	return b.String()
}

func rmr(paths ...string) {
	for _, path := range paths {
		if debug {
			log.Println("rm -r", path)
		}
		os.RemoveAll(path)
	}
}

func getReleaseVersion() (string, error) {
	bs, err := os.ReadFile("RELEASE")
	if err != nil {
		return "", err
	}
	return string(bytes.TrimSpace(bs)), nil
}

func getGitVersion() (string, error) {
	// The current version as Git sees it
	bs, err := runError("git", "describe", "--always", "--dirty", "--abbrev=8")
	if err != nil {
		return "", err
	}
	vcur := string(bs)

	// The closest current tag name
	bs, err = runError("git", "describe", "--always", "--abbrev=0")
	if err != nil {
		return "", err
	}
	v0 := string(bs)

	// To be more semantic-versionish and ensure proper ordering in our
	// upgrade process, we make sure there's only one hyphen in the version.

	versionRe := regexp.MustCompile(`-([0-9]{1,3}-g[0-9a-f]{5,10}(-dirty)?)`)
	if m := versionRe.FindStringSubmatch(vcur); len(m) > 0 {
		suffix := strings.ReplaceAll(m[1], "-", ".")

		if strings.Contains(v0, "-") {
			// We're based of a tag with a prerelease string. We can just
			// add our dev stuff directly.
			return fmt.Sprintf("%s.dev.%s", v0, suffix), nil
		}

		// We're based on a release version. We need to bump the patch
		// version and then add a -dev prerelease string.
		next := nextPatchVersion(v0)
		return fmt.Sprintf("%s-dev.%s", next, suffix), nil
	}
	return vcur, nil
}

func getVersion() string {
	// First try for a RELEASE file,
	if ver, err := getReleaseVersion(); err == nil {
		return ver
	}
	// ... then see if we have a Git tag.
	if ver, err := getGitVersion(); err == nil {
		if strings.Contains(ver, "-") {
			// The version already contains a hash and stuff. See if we can
			// find a current branch name to tack onto it as well.
			return ver + getBranchSuffix()
		}
		return ver
	}
	// This seems to be a dev build.
	return "unknown-dev"
}

func semanticVersion() (major, minor, patch int) {
	r := regexp.MustCompile(`v(\d+)\.(\d+).(\d+)`)
	matches := r.FindStringSubmatch(getVersion())
	if len(matches) != 4 {
		return 0, 0, 0
	}

	var ints [3]int
	for i, s := range matches[1:] {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints[0], ints[1], ints[2]
}

func getBranchSuffix() string {
	bs, err := runError("git", "branch", "-a", "--contains")
	if err != nil {
		return ""
	}

	branches := strings.Split(string(bs), "\n")
	if len(branches) == 0 {
		return ""
	}

	branch := ""
	for i, candidate := range branches {
		if strings.HasPrefix(candidate, "*") {
			// This is the current branch. Select it!
			branch = strings.TrimLeft(candidate, " \t*")
			break
		} else if i == 0 {
			// Otherwise the first branch in the list will do.
			branch = strings.TrimSpace(branch)
		}
	}

	if branch == "" {
		return ""
	}

	// The branch name may be on the form "remotes/origin/foo" from which we
	// just want "foo".
	parts := strings.Split(branch, "/")
	if len(parts) == 0 || len(parts[len(parts)-1]) == 0 {
		return ""
	}

	branch = parts[len(parts)-1]
	switch branch {
	case "master", "release", "main":
		// these are not special
		return ""
	}

	validBranchRe := regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`)
	if !validBranchRe.MatchString(branch) {
		// There's some odd stuff in the branch name. Better skip it.
		return ""
	}

	return "-" + branch
}

func buildStamp() int64 {
	// If SOURCE_DATE_EPOCH is set, use that.
	if s, _ := strconv.ParseInt(os.Getenv("SOURCE_DATE_EPOCH"), 10, 64); s > 0 {
		return s
	}

	// Try to get the timestamp of the latest commit.
	bs, err := runError("git", "show", "-s", "--format=%ct")
	if err != nil {
		// Fall back to "now".
		return time.Now().Unix()
	}

	s, _ := strconv.ParseInt(string(bs), 10, 64)
	return s
}

func buildUser() string {
	if v := os.Getenv("BUILD_USER"); v != "" {
		return v
	}

	u, err := user.Current()
	if err != nil {
		return "unknown-user"
	}
	return strings.Replace(u.Username, " ", "-", -1)
}

func buildHost() string {
	if v := os.Getenv("BUILD_HOST"); v != "" {
		return v
	}

	h, err := os.Hostname()
	if err != nil {
		return "unknown-host"
	}
	return h
}

func buildArch() string {
	os := goos
	if os == "darwin" {
		os = "macos"
	}
	return fmt.Sprintf("%s-%s", os, goarch)
}

func archiveName(target target) string {
	return fmt.Sprintf("%s-%s-%s", target.name, buildArch(), version)
}

func runError(cmd string, args ...string) ([]byte, error) {
	if debug {
		t0 := time.Now()
		log.Println("runError:", cmd, strings.Join(args, " "))
		defer func() {
			log.Println("... in", time.Since(t0))
		}()
	}
	ecmd := exec.Command(cmd, args...)
	bs, err := ecmd.CombinedOutput()
	return bytes.TrimSpace(bs), err
}

func runPrint(cmd string, args ...string) {
	runPrintInDir(".", cmd, args...)
}

func runPrintInDir(dir string, cmd string, args ...string) {
	if debug {
		t0 := time.Now()
		log.Println("runPrint:", cmd, strings.Join(args, " "))
		defer func() {
			log.Println("... in", time.Since(t0))
		}()
	}
	ecmd := exec.Command(cmd, args...)
	ecmd.Stdout = os.Stdout
	ecmd.Stderr = os.Stderr
	ecmd.Dir = dir
	err := ecmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func runPipe(file, cmd string, args ...string) {
	if debug {
		t0 := time.Now()
		log.Println("runPipe:", cmd, strings.Join(args, " "))
		defer func() {
			log.Println("... in", time.Since(t0))
		}()
	}
	fd, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	ecmd := exec.Command(cmd, args...)
	ecmd.Stdout = fd
	ecmd.Stderr = os.Stderr
	err = ecmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fd.Close()
}

func tarGz(out string, files []archiveFile) {
	fd, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}

	gw, err := gzip.NewWriterLevel(fd, gzip.BestCompression)
	if err != nil {
		log.Fatal(err)
	}
	tw := tar.NewWriter(gw)

	for _, f := range files {
		sf, err := os.Open(f.src)
		if err != nil {
			log.Fatal(err)
		}

		info, err := sf.Stat()
		if err != nil {
			log.Fatal(err)
		}
		h := &tar.Header{
			Name:    f.dst,
			Size:    info.Size(),
			Mode:    int64(info.Mode()),
			ModTime: info.ModTime(),
		}

		err = tw.WriteHeader(h)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(tw, sf)
		if err != nil {
			log.Fatal(err)
		}
		sf.Close()
	}

	err = tw.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = gw.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = fd.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func zipFile(out string, files []archiveFile) {
	fd, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}

	zw := zip.NewWriter(fd)

	var fw *flate.Writer

	// Register the deflator.
	zw.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		var err error
		if fw == nil {
			// Creating a flate compressor for every file is
			// expensive, create one and reuse it.
			fw, err = flate.NewWriter(out, flate.BestCompression)
		} else {
			fw.Reset(out)
		}
		return fw, err
	})

	for _, f := range files {
		sf, err := os.Open(f.src)
		if err != nil {
			log.Fatal(err)
		}

		info, err := sf.Stat()
		if err != nil {
			log.Fatal(err)
		}

		fh, err := zip.FileInfoHeader(info)
		if err != nil {
			log.Fatal(err)
		}
		fh.Name = filepath.ToSlash(f.dst)
		fh.Method = zip.Deflate

		if strings.HasSuffix(f.dst, ".txt") {
			// Text file. Read it and convert line endings.
			bs, err := io.ReadAll(sf)
			if err != nil {
				log.Fatal(err)
			}
			bs = bytes.Replace(bs, []byte{'\n'}, []byte{'\r', '\n'}, -1)
			fh.UncompressedSize = uint32(len(bs))
			fh.UncompressedSize64 = uint64(len(bs))

			of, err := zw.CreateHeader(fh)
			if err != nil {
				log.Fatal(err)
			}
			of.Write(bs)
		} else {
			// Binary file. Copy verbatim.
			of, err := zw.CreateHeader(fh)
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.Copy(of, sf)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	err = zw.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = fd.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func codesign(target target) {
	switch goos {
	case "windows":
		windowsCodesign(target.BinaryName())
	case "darwin":
		macosCodesign(target.BinaryName())
	}
}

func macosCodesign(file string) {
	if pass := os.Getenv("CODESIGN_KEYCHAIN_PASS"); pass != "" {
		bs, err := runError("security", "unlock-keychain", "-p", pass)
		if err != nil {
			log.Println("Codesign: unlocking keychain failed:", string(bs))
			return
		}
	}

	if id := os.Getenv("CODESIGN_IDENTITY"); id != "" {
		bs, err := runError("codesign", "--options=runtime", "-s", id, file)
		if err != nil {
			log.Println("Codesign: signing failed:", string(bs))
			return
		}
		log.Println("Codesign: successfully signed", file)
	}
}

func windowsCodesign(file string) {
	st := "signtool.exe"

	if path := os.Getenv("CODESIGN_SIGNTOOL"); path != "" {
		st = path
	}

	for i, algo := range []string{"sha1", "sha256"} {
		args := []string{"sign", "/fd", algo}
		if f := os.Getenv("CODESIGN_CERTIFICATE_FILE"); f != "" {
			args = append(args, "/f", f)
		}
		if p := os.Getenv("CODESIGN_CERTIFICATE_PASSWORD"); p != "" {
			args = append(args, "/p", p)
		}
		if tr := os.Getenv("CODESIGN_TIMESTAMP_SERVER"); tr != "" {
			switch algo {
			case "sha256":
				args = append(args, "/tr", tr, "/td", algo)
			default:
				args = append(args, "/t", tr)
			}
		}
		if i > 0 {
			args = append(args, "/as")
		}
		args = append(args, file)

		bs, err := runError(st, args...)
		if err != nil {
			log.Println("Codesign: signing failed:", string(bs))
			return
		}
		log.Println("Codesign: successfully signed", file, "using", algo)
	}
}

func metalint() {
	lazyRebuildAssets()
	runPrint(goCmd, "test", "-run", "Metalint", "./meta")
}

func metalintShort() {
	lazyRebuildAssets()
	runPrint(goCmd, "test", "-short", "-run", "Metalint", "./meta")
}

func (t target) BinaryName() string {
	if goos == "windows" {
		return t.binaryName + ".exe"
	}
	return t.binaryName
}

func protobufVersion() string {
	bs, err := runError(goCmd, "list", "-f", "{{.Version}}", "-m", "github.com/gogo/protobuf")
	if err != nil {
		log.Fatal("Getting protobuf version:", err)
	}
	return string(bs)
}

func currentAndLatestVersions(n int) ([]string, error) {
	bs, err := runError("git", "tag", "--sort", "taggerdate")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bs), "\n")
	reverseStrings(lines)

	// The one at the head is the latest version. We always keep that one.
	// Then we filter out remaining ones with dashes (pre-releases etc).

	latest := lines[:1]
	nonPres := filterStrings(lines[1:], func(s string) bool { return !strings.Contains(s, "-") })
	vers := append(latest, nonPres...)
	return vers[:n], nil
}

func reverseStrings(ss []string) {
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[len(ss)-1-i] = ss[len(ss)-1-i], ss[i]
	}
}

func filterStrings(ss []string, op func(string) bool) []string {
	n := ss[:0]
	for _, s := range ss {
		if op(s) {
			n = append(n, s)
		}
	}
	return n
}

func tagMessage(tag string) (string, error) {
	hash, err := runError("git", "rev-parse", tag)
	if err != nil {
		return "", err
	}
	obj, err := runError("git", "cat-file", "-p", string(hash))
	if err != nil {
		return "", err
	}
	return trimTagMessage(string(obj), tag), nil
}

func trimTagMessage(msg, tag string) string {
	firstBlank := strings.Index(msg, "\n\n")
	if firstBlank > 0 {
		msg = msg[firstBlank+2:]
	}
	msg = strings.TrimPrefix(msg, tag)
	beginSig := strings.Index(msg, "-----BEGIN PGP")
	if beginSig > 0 {
		msg = msg[:beginSig]
	}
	return strings.TrimSpace(msg)
}

func nextPatchVersion(ver string) string {
	parts := strings.SplitN(ver, "-", 2)
	digits := strings.Split(parts[0], ".")
	n, _ := strconv.Atoi(digits[len(digits)-1])
	digits[len(digits)-1] = strconv.Itoa(n + 1)
	return strings.Join(digits, ".")
}
