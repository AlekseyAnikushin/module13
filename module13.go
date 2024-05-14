package module13

import (
	"encoding/json"
	"encoding/xml"
	"os"
	//"sort"
)

type patient struct {
	Name  string
	Age   int
	Email string
}

type patients struct {
	List []patient `xml:"Patient"`
}

func read(inputFile string, res *[]patient) error {
	f, err := os.Open(inputFile)

	if err != nil {
		return err
	}

	defer f.Close()

	var p patient
	dec := json.NewDecoder(f)
	*res = make([]patient, 0, 3)
	for dec.More() {
		err = dec.Decode(&p)
		if err != nil {
			return err
		}
		*res = append(*res, p)
	}

	return nil
}

func write(outputFile string, data *patients) error {
	f, err := os.Create(outputFile)

	if err != nil {
		return err
	}

	defer f.Close()

	f.WriteString(xml.Header)

	enc := xml.NewEncoder(f)
	enc.Indent("", "  ")
	err = enc.Encode(data)

	return err
}

func Do(inFile string, outFile string) error {
	var p []patient

	err := read(inFile, &p)

	if err != nil {
		return err
	}

	//sort.SliceStable(p, func(i, j int) bool {
	//	return p[i].Age < p[j].Age
	//})

	if p != nil {
		ps := patients{p}
		err = write(outFile, &ps)
		if err != nil {
			return err
		}
	}

	return nil
}
