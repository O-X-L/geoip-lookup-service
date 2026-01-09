package u

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"git.oxl.at/geoip-lookup-service/internal/cnf"
)

func checkGeoIPDB(file string) error {
	if file != "" {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return fmt.Errorf("provided db-file '%s' does not exist - %v", file, err)
		}
	}
	return nil
}

func CheckGeoIPDBs() error {
	if err := checkGeoIPDB(cnf.DB_LITE); err != nil {
		return err
	}
	if err := checkGeoIPDB(cnf.DB_COUNTRY); err != nil {
		return err
	}
	if err := checkGeoIPDB(cnf.DB_CITY); err != nil {
		return err
	}
	if err := checkGeoIPDB(cnf.DB_ASN); err != nil {
		return err
	}
	if err := checkGeoIPDB(cnf.DB_PRIVACY); err != nil {
		return err
	}
	return nil
}

func prependDBDirPath(dirPath string, filePath string) string {
	if filePath == "" || strings.Contains(filePath, "/") {
		return filePath
	}
	return filepath.Join(dirPath, filePath)
}

func PrependDBDirPaths() {
	dbDir := os.Getenv(cnf.ENV_PATH_DB)
	if dbDir == "" {
		return
	}
	cnf.DB_LITE = prependDBDirPath(dbDir, cnf.DB_LITE)
	cnf.DB_COUNTRY = prependDBDirPath(dbDir, cnf.DB_COUNTRY)
	cnf.DB_CITY = prependDBDirPath(dbDir, cnf.DB_CITY)
	cnf.DB_ASN = prependDBDirPath(dbDir, cnf.DB_ASN)
	cnf.DB_PRIVACY = prependDBDirPath(dbDir, cnf.DB_PRIVACY)
}
