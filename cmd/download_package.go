package cmd

import (
	"os"
	"path"

	"github.com/nanovms/ops/lepton"
	api "github.com/nanovms/ops/lepton"
	"github.com/nanovms/ops/log"
	"github.com/nanovms/ops/types"
)

func downloadLocalPackage(pkg string) string {
	packagesDirPath := path.Join(api.GetOpsHome(), "local_packages")
	return downloadAndExtractPackage(packagesDirPath, pkg, lepton.NewConfig())
}

func packageDirectoryPath() string {
	return path.Join(api.GetOpsHome(), "packages")
}

func downloadPackage(pkg string, config *types.Config) string {
	return downloadAndExtractPackage(packageDirectoryPath(), pkg, config)
}

func downloadAndExtractPackage(packagesDirPath, pkg string, config *types.Config) string {
	err := os.MkdirAll(packagesDirPath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	expackage := path.Join(packagesDirPath, pkg)
	opsPackage, err := api.DownloadPackage(pkg, config)
	if err != nil {
		log.Fatal(err)
	}

	api.ExtractPackage(opsPackage, packagesDirPath, config)

	err = os.Remove(opsPackage)
	if err != nil {
		log.Fatal(err)
	}

	return expackage
}
