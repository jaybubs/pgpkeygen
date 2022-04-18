package main

import (
	"fmt"
	"os"
	"text/template"

	. "pgpkeygen/utilities"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
)

type Template_vars struct {
	Username	string
	Email		string
	Fingerprint	string
}

func Parse_and_save(template_path string, output_path string, tvars *Template_vars) error {

	tmpl, err := template.ParseFiles(template_path)
	CE(err)
	// save to file
	var output_file *os.File
	output_file, err = os.Create(output_path)
	tmpl.Execute(output_file, tvars)
	// and back to string for db.exec

	return nil

}

func main() {

	var user_info Template_vars


	user_info.Username = os.Getenv("NUSER")
	user_info.Email = os.Getenv("NMAIL")

	mykey, err := crypto.GenerateKey(user_info.Username,user_info.Email,"RSA",4096)
	CE(err)
	bla, err := mykey.GetArmoredPublicKeyWithCustomHeaders("","")
	CE(err)
	fmt.Println(bla)
	bleh, err := mykey.ArmorWithCustomHeaders("","")
	CE(err)
	fmt.Println(bleh)

	err = os.WriteFile("./publickey", []byte(bla), 0644)
	CE(err)

	user_info.Fingerprint = mykey.GetFingerprint()
	err = os.WriteFile("./fingerprint", []byte(user_info.Fingerprint), 0644)
	
	
	Parse_and_save("./gitignore.tmpl", "./gitconfig", &user_info)
	
	myring, err := crypto.NewKeyRing(mykey)
	CE(err)
	var kbx *os.File
	kbx, err = os.Create("./mykeyring.kbx")
	err = os.WriteFile(kbx, []byte(myring), 0644)
}
