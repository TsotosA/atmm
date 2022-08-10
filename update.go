package main

import (
	"fmt"
	"github.com/tsotosa/atmm/config"
	"github.com/tsotosa/atmm/gconst"
	"github.com/tsotosa/atmm/helper"
	"github.com/tsotosa/atmm/model"
	"go.uber.org/zap"
	"os"
	"runtime"
	"sort"
	"strings"
)

func HandleUpdate() error {
	zap.S().Debugf("HandleUpdate()")
	allReleases, err := GetGithubReleases(1)
	if err != nil {
		zap.S().Warnf("could not get release list from github with err: %v", err)
		return err
	}
	var releases []model.GithubRelease
	for i, _ := range allReleases {
		if !allReleases[i].Draft {
			releases = append(releases, allReleases[i])
		}
	}
	if releases == nil || (len(releases) <= 0) {
		zap.S().Infof("no releases found")
		return nil
	}
	sort.Slice(releases, func(i, j int) bool {
		return releases[i].CreatedAt > releases[j].CreatedAt
	})
	latestVersion := releases[0].TagName
	isSameVersion := latestVersion == Version
	if isSameVersion {
		zap.S().Debugf("already at latest version: %s", Version)
		return nil
	}
	latestVersionAssets := releases[0].Assets
	var latestApplicableAsset model.GithubReleaseAsset
	for _, asset := range latestVersionAssets {
		if strings.Contains(asset.Name, runtime.GOOS) {
			zap.S().Infof("found matching new release asset to use for update: [%s]", asset.Name)
			latestApplicableAsset = asset
		}
	}
	if latestApplicableAsset.Id == 0 {
		zap.S().Infof("could not locate an applicable asset in release with GOOS: %s", runtime.GOOS)
		return nil
	}
	err = DownloadUrlToLocation(latestApplicableAsset.Name, latestApplicableAsset.Url, "./tmp/")
	defer func() {
		err := os.RemoveAll(gconst.TmpDir)
		if err != nil {
			zap.S().Warnf("failed to remove tmp directory")
		}
	}()
	if err != nil {
		zap.S().Warnf("could not download binary to tmp location with error: %v", err)
		return err
	}
	path, err := helper.CurrrentBinaryAbsolutePath()
	if err != nil {
		zap.S().Warnf("could not get current binary path with error: %v", err)
		return err
	}
	err = os.Rename(path, fmt.Sprintf("%s%s", path, ".backup"))
	if err != nil {
		zap.S().Warnf("could not backup original binary with error: %v", err)
		return err
	}
	err = os.Rename(fmt.Sprintf("%s%s", "./tmp/", latestApplicableAsset.Name), path)
	if err != nil {
		zap.S().Warnf("could not replace original binary with updated with error: %v", err)
		return err
	}
	if config.Conf.IsAutoRestartManaged {
		zap.S().Infof("IsAutoRestartManaged flag is true, exiting self and hoping somebody else will resurrect me :-)")
		panic(Exit{3}) // 3 is the exit code
	}
	zap.S().Infof("IsAutoRestartManaged flag is false, update was sucesfull but application restart is needed")
	return nil
}
