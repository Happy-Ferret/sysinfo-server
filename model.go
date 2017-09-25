package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"strings"
)

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS servers
(machineid TEXT NOT NULL PRIMARY KEY,
hostname TEXT NOT NULL,
data jsonb
)`

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func getMachineIDs(db *sql.DB) ([]MachineID, error) {
	mm := []MachineID{}

	rows, err := db.Query(
		"SELECT machineid,hostname FROM servers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m MachineID
		if err := rows.Scan(&m.MachineID, &m.Hostname); err != nil {
			return nil, err
		}
		mm = append(mm, m)
	}

	return mm, nil
}

func getPackagesByMI(db *sql.DB, machineid string) (Packages, error) {
	var p Packages
	var ss []string

	rows, err := db.Query(
		"SELECT value FROM servers, jsonb_array_elements((servers.data->>'packages')::jsonb) WHERE machineid = $1", machineid)
	if err != nil {
		return p, err
	}
	defer rows.Close()

	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			return p, err
		}
		s = strings.Trim(s, `"`)
		ss = append(ss, s)
	}

	p.Packages = ss
	return p, nil
}

func getSysInfoByMI(db *sql.DB, machineid string) (SysInfo, error) {
	var si SysInfo
	var s string
	var err error

	row := db.QueryRow(
		"SELECT data FROM servers WHERE machineid = $1", machineid)
	err = row.Scan(&s)
	if err != nil {
		return si, err
	}
	err = json.Unmarshal([]byte(s), &si)
	if err != nil {
		return si, err
	}
	return si, nil
}

func writeDataToDB(db *sql.DB, msg []byte) error {
	var si SysInfo
	var err error

	err = json.Unmarshal(msg, &si)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`INSERT INTO servers VALUES ($1, $2, $3)
    ON CONFLICT (machineid) DO UPDATE SET
    hostname = $2,
    data = $3`, si.Node.MachineID, si.Node.Hostname, string(msg))
	if err != nil {
		return err
	}

	return nil
}
