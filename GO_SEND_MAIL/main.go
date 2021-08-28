package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"
)

func main() {
	subjectMenssa := "Mi primer correo enviado desde GO"
	mailTarget := MailTarget{}
	mailTarget.init("blxxxxxxxo@gmail.com", "menssage xxxx xvio", "xxxxxxxr@gmail.com", "Alejandro")
	autentication := AutenticationData{Identity: "", Username: "wxxxxxxxxxx@gmail.com", Password: "password....."}
	SendMail(mailTarget, subjectMenssa, autentication)
}

const (
	HostConst       = "smtp.gmail.com"
	servernameConst = "smtp.gmail.com:465"
)

type Dest struct {
	Name string
}

type MailTarget struct {
	from mail.Address
	to   mail.Address
}

func (m *MailTarget) init(from string, fromName string, to string, toName string) {
	m.from = mail.Address{Name: fromName, Address: from}
	m.to = mail.Address{Name: toName, Address: to}
}

type AutenticationData struct {
	Identity string
	Username string
	Password string
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("error en check error")
		fmt.Println(err)
		log.Panic(err)
	}
}

func SendMail(mailTargets MailTarget, meSubject string, autentication AutenticationData) {
	log.Println("inicia...")
	menssage := ""
	from := mailTargets.from
	to := mailTargets.to

	subject := meSubject
	dest := Dest{Name: to.Address}
	log.Println("inicio de header...")
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	for key, valor := range headers {
		menssage += fmt.Sprintf("%s: %s\r\n", key, valor)
	}

	log.Println("inicio de template...")
	menssa, err := GetTemplate("template.html", dest) // add template.
	log.Println("error template")
	log.Printf("error %s", err)
	checkErr(err)

	log.Println("set del mensaje...")
	menssage += menssa

	log.Println("datos del servidor")
	// datos del servidor
	servername := servernameConst
	host := HostConst

	log.Println("para setear los datos de autenticacion")
	// para setear los datos de autenticacion
	auth := smtp.PlainAuth(autentication.Identity, autentication.Username, autentication.Password, host)

	log.Println("error para setear los datos de autenticacion")
	log.Printf("error %s", err)
	checkErr(err)

	// configuracion de firma de seguridad
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// coneccion con el servidor de mail
	log.Println("inicio la coneccion con el servidor de mail")
	conn, err := tls.Dial("tcp", servername, tlsConfig)
	log.Println("inicio de Error coneccion con el servidor de mail")
	log.Printf("error %s", err)
	checkErr(err)

	// creamos el cliente
	log.Println("inicio creamos el cliente")
	client, err := smtp.NewClient(conn, host)
	log.Println(" error creamos el cliente")
	log.Printf("error %s", err)
	checkErr(err)

	// se la pasa los datos de autenticacion al cliente
	err = client.Auth(auth)
	checkErr(err)

	// se pasa para a quien se le envia el correo
	err = client.Mail(from.Address)
	checkErr(err)

	//se le pada queien envia el correo al cliente
	err = client.Rcpt(to.Address)
	checkErr(err)

	//pasamos la data al cliente
	w, err := client.Data()
	checkErr(err)

	// escribimos el mensaje en el correo
	log.Println(" inicio escribimos el mensaje en el correo")
	_, err = w.Write([]byte(menssage))
	log.Println(" error  escribimos el mensaje en el correo")
	log.Printf("error %s", err)
	checkErr(err)

	// cerramos la escritura
	err = w.Close()
	checkErr(err)

	//cerramos el cliente
	client.Quit()

	log.Print("finalizo el envio del correo exitosamente")

}

// se llama al tempate (templateName = tempate.html)
func GetTemplate(templateName string, dest Dest) (string, error) {
	t, err := template.ParseFiles(templateName) // se llama al tempate
	checkErr(err)
	buf := new(bytes.Buffer)   //se crea un buffer para pasar lo datos al template
	err = t.Execute(buf, dest) // se le pasa los datos al template en bytes
	checkErr(err)
	return buf.String(), err
}

//************************************ old metodo funcional ***********************************************************************************

// func SENDmAIL() {
// 	from := mail.Address{"Bluxxxxxxxxx", "xxxxxxxxx@gmail.com"}
// 	to := mail.Address{"Crxxxxxx Casxxxxx", "xxxxxxry@gmail.com"}
// 	subject := "Mi primer correo enviado desde GO"
// 	dest := Dest{Name: to.Address}

// 	headers := make(map[string]string)
// 	headers["From"] = from.String()
// 	headers["To"] = to.String()
// 	headers["Subject"] = subject
// 	headers["Content-Type"] = `text/html; charset="UTF-8"`

// 	menssage := ""
// 	for key, valor := range headers {
// 		menssage += fmt.Sprintf("%s: %s\r\n", key, valor)
// 	}

// 	// se llama al tempate
// 	t, err := template.ParseFiles("template.html")
// 	checkErr(err)
// 	//se crea un buffer para pasar lo datos al template
// 	buf := new(bytes.Buffer)
// 	err = t.Execute(buf, dest) // se le pasa los datos al template en bytes
// 	checkErr(err)
// 	menssage += buf.String()

// 	// datos del servidor

// 	servername := "smtp.gmail.com:465"
// 	host := "smtp.gmail.com"

// 	// para setear los datos de autenticacion
// 	auth := smtp.PlainAuth("", "bluxxxxxx@gmail.com", "password.....", host)
// 	checkErr(err)

// 	// configuracion de firma de seguridad
// 	tlsConfig := &tls.Config{
// 		InsecureSkipVerify: true,
// 		ServerName:         host,
// 	}

// 	// coneccion con el servidor de mail
// 	conn, err := tls.Dial("tcp", servername, tlsConfig)
// 	checkErr(err)

// 	// creamos el cliente
// 	client, err := smtp.NewClient(conn, host)
// 	checkErr(err)

// 	// se la pasa los datos de autenticacion al cliente
// 	err = client.Auth(auth)
// 	checkErr(err)
// 	// se pasa para a quien se le envia el correo
// 	err = client.Mail(from.Address)
// 	checkErr(err)
// 	//se le pada queien envia el correo al cliente
// 	err = client.Rcpt(to.Address)
// 	checkErr(err)
// 	//pasamos la data al cliente
// 	w, err := client.Data()
// 	checkErr(err)
// 	// escribimos el mensaje en el correo
// 	_, err = w.Write([]byte(menssage))
// 	checkErr(err)
// 	// cerramos la escritura
// 	err = w.Close()
// 	checkErr(err)

// 	//cerramos el cliente
// 	client.Quit()

// 	log.Print("finalizo el envio del correo exitosamente")

// }
