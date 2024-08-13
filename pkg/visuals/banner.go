package visuals

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/projectdiscovery/gologger"
)

func PrintBanner() {
	Version := "2.1.0"

	banner := (`
                                                        :=*#%@@@@#                        
                                    :::::.            =%@@#*++===.                        
                               .=*%@@@@@@@@#+:        :%@@*:                              
                             :*@@@@@@@@@@@@@@@*.        =%@@+                             
                            *@@@@@@@@@@@@@@@@@@%:     .=%@@#:                             
                           #@@@@@@@@@@@@#+=+@@@@%  :=*@@%*++#%%%#*+=:.                    
                          =@@@@@@@@@#+-    =@@@@@%@%#+-:  +%@@%++#%@@#:                   
                          =++**++=:  :--   *@@@@@%#*+====+*#@@@:   :.                     
                          . -      :==-   .%@@@@@=-=*##%%%%#*=.                           
                   :=+=     .=*#*+=-:::::=%@@@@@%                                         
                 -%@@%-     *@@@@@@@@@@@@@@@@@@*..:::-----====---:..                      
                *@@@@@=      =%@@@@@@@@@@@@@@@@%@@@@@@@@@@@@@@@@@@@@%#*+-.                
              .#@@@%+:        .#@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%*:             
              #@@@+        -+#%@@@@@@@@@@@@@@@@@@@@@@@@@@@%%#*******#%%@@@@@@+            
             #@@@@=   :-+#@@@@@@@@@@@@@@@@@@@@@@@@@@@#+=:.             *@@@@%-            
            +@@@@@@@%@@@@@@@@@@@@@@@@@@@@@@@@@@@@%*-                    :--.              
            *@@@@@@@@@@@@@@%@@@@@@@@@@@@@@@@@@@%=                                         
             =*#%%@@@@@@@@%+@@@@@@@@@@@@@@@@@@#                                           
                           .@@@@@@@@@@@@@@@@@%.                                           
                            *@@@@@@@@@@@@@@@@%                                            
                             +@@@@@@@@@@@@@@@@.                                           
                              .+%@@@@@@@@@@@@@#                                           
                                 :=*%@@@@@@@@@@*                                          
                                     .-=*#%@@@@@#.                                        
                                           .:-+*#%-                                       
`)
	banner += PrintRandomColor(`
____,  ____, ____,  __, ,    _,  _, __,   _,  _,  _,    ____,
(-|    (-/_| (-(__  ( |_/    (-|\ | (-|   (-|\ |  (-|   (-/_| 
 _|,   _/  |, ____)  _| \,    _| \|, _|_,  _| \|, __|,  _/  |,
(     (      (      (        (      (     (      (     (      
	`)
	fmt.Fprintln(os.Stderr, strings.Repeat("=", 80))
	gologger.Print().Str(PrintRandomColor(printLine()), PrintRandomColor(Version)).Msg(banner)
	fmt.Fprintln(os.Stderr, strings.Repeat("=", 80))

}

func printLine() string {
	rand.Seed(time.Now().UnixNano())
	lines := []string{
		"Made For 🥷  by Robensive\t\t\t\tVersion",
		"Made With 💝 by Robensive\t\t\t\tVersion",
		"Tasks Automation Framework by Robensive\t\tVersion",
		"The Ultimate Tasks Automation Framework \t\tVersion",
		"Inspired By trickest.io\t\t\t\tVersion",
	}

	randomIndex := rand.Intn(len(lines))

	return lines[randomIndex]
}
