package sqlht

import (
	"database/sql"
	"errorssx"
)

const rMyName = "sqlht"

type WhoMe struct{}

func (whoMe *WhoMe) Name() string {
	return rMyName

}

func Query(db *sql.DB, query string, args ...interface{}) ([]map[string]string, error) {
	sqlRows, err := db.Query(query, args...)
	if err != nil {
		return nil, errorssx.New(&WhoMe{}, "Query: querying database", err)
	}

	columns, err := sqlRows.Columns()
	if err != nil {
		return nil, errorssx.New(&WhoMe{}, "Query: retrieving columns", err)
	}

	allData := make([]map[string]string, 0)
	var scanners = make([]string, len(columns))
	var scannerPs = make([]interface{}, len(columns))
	for indie := range scannerPs {
		scannerPs[indie] = &scanners[indie]
	}

	for sqlRows.Next() {
		err := sqlRows.Scan(scannerPs...)
		if err != nil {
			return nil, errorssx.New(&WhoMe{}, "Query: scanning data", err)
		}
		data := make(map[string]string)
		for indie := range scanners {
			data[columns[indie]] = scanners[indie]
		}
		allData = append(allData, data)
	}

	return allData, nil
}
