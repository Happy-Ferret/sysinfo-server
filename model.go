package main

import (
	"database/sql"
	"encoding/json"
	"strings"
)

/*
CREATE TABLE IF NOT EXISTS servers
(id SERIAL,
machineid TEXT NOT NULL,
hostname TEXT NOT NULL,
data jsonb
);
*/

func getMachineIDs(db *sql.DB) ([]MachineID, error) {
	mm := []MachineID{}

	rows, err := db.Query("SELECT machineid,hostname FROM servers")
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
