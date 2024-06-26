package module13

import (
	"encoding/json"
	"os"
	"sort"
)

type Patient struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func read(inputFile string, res *[]Patient) error {
	f, err := os.Open(inputFile)

	if err != nil {
		return err
	}

	defer f.Close()

	var p Patient
	dec := json.NewDecoder(f)
	*res = make([]Patient, 0, 3)
	for dec.More() {
		err = dec.Decode(&p)
		if err != nil {
			return err
		}
		*res = append(*res, p)
	}

	return nil
}

func write(outputFile string, data *[]Patient) error {
	f, err := os.Create(outputFile)

	if err != nil {
		return err
	}

	defer f.Close()

	enc := json.NewEncoder(f)
	err = enc.Encode(data)

	return err
}

func Do(inFile string, outFile string) error {
	var p []Patient

	err := read(inFile, &p)

	if err != nil {
		return err
	}

	sort.SliceStable(p, func(i, j int) bool {
		return p[i].Age < p[j].Age
	})

	if p != nil {
		err = write(outFile, &p)
		if err != nil {
			return err
		}
	}

	return nil
}
