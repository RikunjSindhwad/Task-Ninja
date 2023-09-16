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

- :robot: **Task Automation**: Define and automate tasks, creating a seamless workflow for your projects.

- :rocket: **Parallel Execution**: Execute tasks in parallel to save time and improve efficiency.

- :link: **Dependency Management**: Specify task dependencies to ensure that tasks are executed in the correct order.

- :wrench: **Customization**: Tailor TaskNinja to your specific needs by defining custom tasks and workflows.
  
- :scroll: **Logging**: TaskNinja provides comprehensive logging capabilities to help you monitor and troubleshoot your automation processes.
  
- :jigsaw: **Extensibility**: Easily add new functionalities and plugins to TaskNinja for enhanced capabilities.


## Acknowledgments

TaskNinja is inspired by the amazing work of the [trickest.io](https://trickest.io) platform and the [Raydar](https://github.com/devanshbatham/rayder) tool. We extend our gratitude to their contributions to the task automation and cybersecurity community.

## Sample - Static 
```bash
 root@robensive>  Task-Ninja -w zap_scanner.yaml -nb -v url_list=urls.txt
================================================================================
[Workflow-Credit] Tasked Workflow 'ZAP Scanner' Workflow-Author=Rikunj Sindhwad
================================================================================
------------------------------------------------------------------------------------------------------------------------
[START] [2023-09-16T09:12:36Z] Task Started TaskName=Create Required Directories
------------------------------------------------------------------------------------------------------------------------
[Task-Info] Task is Static TaskName=Create Required Directories
[Static-Task: Create Required Directories] [2023-09-16T09:12:36Z] Executing Task
------------------------------------------------------------------------------------------------------------------------
[SUCCESS] [2023-09-16T09:12:36Z] Task Finished TaskName=Create Required Directories
------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------------
[START] [2023-09-16T09:12:37Z] Task Started TaskName=Downnload zap yaml
------------------------------------------------------------------------------------------------------------------------
[Task-Info] Task is Static TaskName=Downnload zap yaml
[Static-Task: Downnload zap yaml] [2023-09-16T09:12:37Z] Executing Task
------------------------------------------------------------------------------------------------------------------------
[SUCCESS] [2023-09-16T09:12:38Z] Task Finished TaskName=Downnload zap yaml
------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------------
[START] [2023-09-16T09:12:39Z] Task Started TaskName=Modify ZAP config
------------------------------------------------------------------------------------------------------------------------
[Task-Info] Task is Static TaskName=Modify ZAP config
[Static-Task: Modify ZAP config] [2023-09-16T09:12:39Z] Executing Task
------------------------------------------------------------------------------------------------------------------------
[SUCCESS] [2023-09-16T09:12:39Z] Task Finished TaskName=Modify ZAP config
------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------------
[START] [2023-09-16T09:12:40Z] Task Started TaskName=run ZAP
------------------------------------------------------------------------------------------------------------------------
[Task-Info] Task is Static TaskName=run ZAP
[Static-Task: run ZAP] [2023-09-16T09:12:40Z] Executing Task
------------------------------------------------------------------------------------------------------------------------
[START] [2023-09-16T09:14:20Z] Task Started TaskName=Result-Check
------------------------------------------------------------------------------------------------------------------------
[Task-Info] Task is Static TaskName=Result-Check
[Static-Task: Result-Check] [2023-09-16T09:14:20Z] Executing Task
"Cross Site Scripting (Reflected)","High (Medium)","https://ginandjuice.shop/catalog/filter?category=Accessories%0A%0D%0A%0D%3CscrIpt%3Ealert%281%29%3B%3C%2FscRipt%3E","category","GET","<scrIpt>alert(1);</scRipt>"
"Cross Site Scripting (Reflected)","High (Medium)","https://ginandjuice.shop/login","username","POST","';alert(1);'"
"Vulnerable JS Library","Medium (Medium)","https://ginandjuice.shop/resources/js/angular_1-7-7.js","","GET","/*
 AngularJS v1.7.7"
result saved in hive/out/ZAP-Result.CSV
------------------------------------------------------------------------------------------------------------------------
[SUCCESS] [2023-09-16T09:14:20Z] Task Finished TaskName=Result-Check
------------------------------------------------------------------------------------------------------------------------
[Workflow-Complete] Workflow 'ZAP Scanner' Execution Complete Workflow-Author=Rikunj Sindhwad
================================================================================
```

### Installation

TaskNinja is easy to install using the following commands:

- [ ] **Go Install**
  ```bash
  GO111MODULE=on
  go install github.com/RikunjSindhwad/Task-Ninja@latest
  ```
- [ ] Build
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


