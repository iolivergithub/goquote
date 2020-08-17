package main  
 
import ( 
	       "fmt"
	       "flag"
	       "os"

	       "github.com/google/go-tpm/tpm2"
       )

var (

    )

func main() {

   var tpmname = flag.String("tpm","/dev/tpm0", "The path to the TPM device")
   flag.Parse()

   rwc,err := tpm2.OpenTPM(*tpmname)
   if err != nil {
          fmt.Fprintf(os.Stderr,"Opening TPM device failed: ",err)
          os.Exit(1)
      }

   defer rwc.Close()


}