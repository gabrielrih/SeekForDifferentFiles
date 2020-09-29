package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	/** Verify the arguments */
	var newDir string
	var oldDir string

	newDir, oldDir = verifyArguments()

	/** Check diretories */
	newFilesArray := make([]string, 0)
	obsoleteFilesArray := make([]string, 0)
	newFilesArray, obsoleteFilesArray = checkDirectories(newDir, oldDir)

	/** Mostra resultado ao usuário */
	fmt.Println("Arquivos novos encontrados:")
	printElementsFromSlice(newFilesArray)

	fmt.Println()
	fmt.Println("Arquivos obsoletos:")
	printElementsFromSlice(obsoleteFilesArray)
}

func verifyArguments() (string, string) {

	oldDirPtr := flag.String("oldDir", "", "Old directory used to compare.")
	newDirPtr := flag.String("newDir", "", "New directory used to compare.")

	flag.Parse()

	if (*oldDirPtr == "") || (*newDirPtr == "") {
		os.Exit(1)
	}

	return string(*newDirPtr), string(*oldDirPtr)
}

func checkDirectories(newDir string, oldDir string) ([]string, []string) {

	/** Returned arrays */
	newFilesArray := make([]string, 0)
	obsoleteFilesArray := make([]string, 0)

	/** Verifica os arquivos contidos no diretório antigo */
	filesOld, err := ioutil.ReadDir(oldDir)
	if err != nil {

		log.Fatal(err)

	} else {

		/** Verifica os arquivos contidos no diretório novo */
		filesNew, err := ioutil.ReadDir(newDir)
		if err != nil {

			log.Fatal(err)

		} else {

			/** Busca os arquivos novos */
			// newFilesArray := make([]string, 0)
			for _, auxFileNew := range filesNew {

				isFound := false
				for _, auxFileOld := range filesOld {
					if auxFileNew.Name() == auxFileOld.Name() {
						isFound = true
						break
					}
				}

				if isFound == false {
					newFilesArray = append(newFilesArray, auxFileNew.Name())
				}

			}

			/** Busca os arquivos obsoletos (existiam no diretório antigo mas nâo existem no novo) */
			// obsoleteFilesArray := make([]string, 0)
			for _, auxFileOld := range filesOld {

				isFound := false
				for _, auxFileNew := range filesNew {
					if auxFileOld.Name() == auxFileNew.Name() {
						isFound = true
						break
					}
				}

				if isFound == false {
					obsoleteFilesArray = append(obsoleteFilesArray, auxFileOld.Name())
				}

			}

		}

	}

	return newFilesArray, obsoleteFilesArray

}

func printElementsFromSlice(s []string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
