package main

import (
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


	user_info.Username = os.Getenv("OIDC_USER")
	user_info.Email = os.Getenv("OIDC_MAIL")

	// standard git requirement, no pass
	gen_key, err := crypto.GenerateKey(user_info.Username,user_info.Email,"RSA",4096)
	CE(err)

	// save pubkey in friendly copypastable format
	pub_key, err := gen_key.GetArmoredPublicKeyWithCustomHeaders("","")
	CE(err)
	err = os.WriteFile("./publickey", []byte(pub_key), 0644)
	CE(err)

	user_info.Fingerprint = gen_key.GetFingerprint()
	
	Parse_and_save("./gitconfig.tmpl", "./.gitconfig", &user_info)
	
	
}
