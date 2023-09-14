# TaskNinja

TaskNinja is a versatile and extensible task automation framework designed to simplify and streamline your workflow. Whether you're managing complex tasks, automating routine operations, or orchestrating a series of commands, TaskNinja is here to make your life easier.


```
                                                                            
                                                              .:--                        
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
                                                  ..                                      
                                                                   
 ____,   ____, ____,  __, ,    _,  _, __,   _,  _,  _,    ____,
(-|     (-/_| (-(__  ( |_/    (-|\ | (-|   (-|\ |  (-|   (-/_| 
 _|,    _/  |, ____)  _| \,    _| \|, _|_,  _| \|, __|,  _/  |,
(      (      (      (        (      (     (      (     (      
         Made For 🥷  by Robensive                               Version=1.0
================================================================================
```


## Features

- **Task Automation**: Define and automate tasks, creating a seamless workflow for your projects.

- **Parallel Execution**: Execute tasks in parallel to save time and improve efficiency.

- **Dependency Management**: Specify task dependencies to ensure that tasks are executed in the correct order.

- **Customization**: Tailor TaskNinja to your specific needs by defining custom tasks and workflows.
  
- **Logging**: TaskNinja provides comprehensive logging capabilities to help you monitor and troubleshoot your automation processes.
  
- **Extensibility**: Easily add new functionalities and plugins to TaskNinja for enhanced capabilities.



## Acknowledgments

TaskNinja is inspired by the amazing work of the [trickest.io](https://trickest.io) platform and the [Raydar](https://github.com/devanshbatham/rayder) tool. We extend our gratitude to their contributions to the task automation and cybersecurity community.

## Sample - Static 
```bash
 root@robensive> Task-Ninja -w zap_scanner.yaml -noBanner -v url_list=urls.txt
[Workflow-Credit] Tasked Workflow 'ZAP Scanner' Workflow-Author=Rikunj Sindhwad
------------------------------------------------------------------------------------------------------------------------
[Start] [2023-09-14T17:45:45Z] Task Started TaskName=Create Required Directories
------------------------------------------------------------------------------------------------------------------------
[Task-Info] Task is Static TaskName=Create Required Directories
[Static-Task: Create Required Directories] [2023-09-14T17:45:45Z] Executing Task
------------------------------------------------------------------------------------------------------------------------
[Success] [2023-09-14T17:45:45Z] Task Finished TaskName=Create Required Directories
------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------------
[Start] [2023-09-14T17:45:46Z] Task Started TaskName=Downnload zap yaml
------------------------------------------------------------------------------------------------------------------------
[Task-Info] Task is Static TaskName=Downnload zap yaml
[Static-Task: Downnload zap yaml] [2023-09-14T17:45:46Z] Executing Task
------------------------------------------------------------------------------------------------------------------------
[Success] [2023-09-14T17:45:46Z] Task Finished TaskName=Downnload zap yaml
------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------------
[Start] [2023-09-14T17:45:47Z] Task Started TaskName=Modify ZAP config
------------------------------------------------------------------------------------------------------------------------
[Task-Info] Task is Static TaskName=Modify ZAP config
[Static-Task: Modify ZAP config] [2023-09-14T17:45:47Z] Executing Task
------------------------------------------------------------------------------------------------------------------------
[Success] [2023-09-14T17:45:52Z] Task Finished TaskName=Modify ZAP config
------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------------
[Start] [2023-09-14T17:45:53Z] Task Started TaskName=run ZAP
------------------------------------------------------------------------------------------------------------------------
[Task-Info] Task is Static TaskName=run ZAP
[Static-Task: run ZAP] [2023-09-14T17:45:53Z] Executing Task
------------------------------------------------------------------------------------------------------------------------
[Success] [2023-09-14T17:46:09Z] Task Finished TaskName=run ZAP
------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------------
[Start] [2023-09-14T17:46:10Z] Task Started TaskName=Result-Check
------------------------------------------------------------------------------------------------------------------------
[Task-Info] Task is Static TaskName=Result-Check
[Static-Task: Result-Check] [2023-09-14T17:46:10Z] Executing Task
result saved in hive/out/ZAP-Result.CSV
------------------------------------------------------------------------------------------------------------------------
[Success] [2023-09-14T17:46:10Z] Task Finished TaskName=Result-Check
------------------------------------------------------------------------------------------------------------------------
[Workflow-Complete] Workflow 'ZAP Scanner' Execution Complete Workflow-Author=Rikunj Sindhwad
```

### Installation

TaskNinja is easy to install using the following commands:

[-] Go Install
```
GO111MODULE=on
go install github.com/RikunjSindhwad/Task-Ninja@latest
```
[-] Build
```bash
# Clone the repository
git clone https://github.com/RikunjSindhwad/Task-Ninja.git

# Navigate to the TaskNinja directory
cd Task-Ninja

# Build TaskNinja
go build
```
## Getting Started

To get started with TaskNinja, please refer to the [Getting Started Guide](https://github.com/RikunjSindhwad/Task-Ninja/wiki/Getting-Started) in our Wiki. It provides step-by-step instructions on installation and usage.

## Contributing

We welcome contributions from the community! If you'd like to contribute to TaskNinja, please follow our [Contribution Guidelines](CONTRIBUTING.md).

## License

TaskNinja is open-source and released under the [MIT License](LICENSE).

## Contact

If you have any questions, feedback, or need assistance, feel free to me on [linkedin](https://www.linkedin.com/in/rikunj/) or [telegram](https://t.me/R0B077)

## Join Robensive Community
Robensive community helps you to get regular updates on news/jobs/tools/courses and knowledge sharing
Join over [telegram](https://t.me/robensive)


