<h1 align="center">
  Task-Ninja
</h1>
<p align="center">
  <strong>The ultimate tasks automation framework for DevSecOps, Hackers, Bugbounty Hunters!</strong>
</p>
<p align="center">
<a href="https://opensource.org/licenses/GPL"><img src="https://img.shields.io/badge/license-GPL-_red.svg"></a>
<a href="https://goreportcard.com/badge/github.com/RikunjSindhwad/Task-Ninja"><img src="https://goreportcard.com/badge/github.com/RikunjSindhwad/Task-Ninja"></a>
<a href="https://github.com/RikunjSindhwad/Task-Ninja/release"><img src="https://img.shields.io/github/release/RikunjSindhwad/Task-Ninja"></a>
<a href="https://twitter.com/sindhwadrikunj"><img src="https://img.shields.io/twitter/follow/sindhwadrikunj.svg?logo=twitter"></a>
</p>

<p align="center">
  <a href="#Features">Features</a> •
  <a href="#Acknowledgments">Acknowledgments</a> •
  <a href="#Samples">Samples</a> •
  <a href="#Installation">Installation</a> •
  <a href="#Getting Started">Getting Started</a> •
  <a href="https://t.me/robensive">Join Telegram</a>
</p>

---
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
         Made For 🥷  by Robensive                               Version=2.0
================================================================================
```

TaskNinja is a versatile and extensible task automation framework designed to simplify and streamline your workflow. Whether you're managing complex tasks, automating routine operations, or orchestrating a series of commands, TaskNinja is here to make your life easier. Task ninja takes input from basic yaml file and executes within docker containers. It has ability to take output from that container and pass it to the next container by just specifying basic yaml input. The purpose to make this tool was to automate some tasks for my personal use. 

> ***Feel free to reachout to me for a private workflows that I have built or for custom workflow for your company needs. This can save your $$$$***


![image](https://github.com/RikunjSindhwad/Task-Ninja/assets/54503807/e4a20fc0-af73-43f1-852d-ab895d8a6de7)


## Features

- :robot: **Task Automation**: Define and automate tasks, creating a seamless workflow for your projects.

- :rocket: **Parallel Execution**: Execute tasks in parallel to save time and improve efficiency.

- :link: **Dependency Management**: Specify task dependencies to ensure that tasks are executed in the correct order.

- :wrench: **Customization**: Tailor TaskNinja to your specific needs by defining custom tasks and workflows.
  
- :scroll: **Logging**: TaskNinja provides comprehensive logging capabilities to help you monitor and troubleshoot your automation processes.
  
- :jigsaw: **Extensibility**: Easily add new functionalities and plugins to TaskNinja for enhanced capabilities.


## Acknowledgments

TaskNinja is inspired by the amazing work of the [trickest.io](https://trickest.io) platform and the [Raydar](https://github.com/devanshbatham/rayder) tool. I extend my gratitude to their contributions to the task automation and cybersecurity community.

## Samples

<details>
<summary>:rocket: Sample - Dynamic & parallel Execution</summary>

### Workflow URL: [passiveSubdomains.yaml](https://github.com/RikunjSindhwad/Task-Ninja-Workflows/blob/main/EASM/Discovery/Subdomains/passiveSubdomains.yaml)
 
```bash
root@robensive> Task-Ninja -nb -w passiveSubdomains.yaml -v tld_list=test/tld.txt
================================================================================
[Workflow-Credit] Tasked Workflow 'Discovery Passive Subdomains' Workflow-Author=Rikunj Sindhwad
================================================================================
[START] [2023-12-04T14:13:31-05:00] Task Started TaskName=Subdomain Center Subdomains
[START] [2023-12-04T14:13:31-05:00] Task Started TaskName=JIDC Subdomains
[SUCCESS] [2023-12-04T14:13:35-05:00] Task Finished TaskName=JIDC Subdomains
[SUCCESS] [2023-12-04T14:13:36-05:00] Task Finished TaskName=Subdomain Center Subdomains
[START] [2023-12-04T14:13:37-05:00] Task Started TaskName=AssetFinder
[SUCCESS] [2023-12-04T14:13:44-05:00] Task Finished TaskName=AssetFinder
[START] [2023-12-04T14:13:45-05:00] Task Started TaskName=Findomain
[SUCCESS] [2023-12-04T14:13:46-05:00] Task Finished TaskName=Findomain
[START] [2023-12-04T14:13:47-05:00] Task Started TaskName=Subfinder
[START] [2023-12-04T14:13:47-05:00] Task Started TaskName=VITA Subdomains
[SUCCESS] [2023-12-04T14:14:09-05:00] Task Finished TaskName=Subfinder
[SUCCESS] [2023-12-04T14:14:21-05:00] Task Finished TaskName=VITA Subdomains
[START] [2023-12-04T14:14:22-05:00] Task Started TaskName=Merge All
[SUCCESS] [2023-12-04T14:14:22-05:00] Task Finished TaskName=Merge All
[START] [2023-12-04T14:14:23-05:00] Task Started TaskName=Result
Check-Result: /home/kali/Workflows/hive/Result/out/result.txt
IdentifiedLines:3
[SUCCESS] [2023-12-04T14:14:24-05:00] Task Finished TaskName=Result
================================================================================
[Workflow-Complete] Workflow 'Discovery Passive Subdomains' Execution Complete Time Taken=53.77675225s
================================================================================

```
</details>

<details>
<summary>:robot: Sample - Static Execution</summary>

### Workflow URL: [apk_url.yaml](https://github.com/RikunjSindhwad/Task-Ninja-Workflows/blob/main/Scanning/apk_url.yaml)

> `-detailed` flag gives more output
 
  ```bash
  root@robensive> Task-Ninja -nb -w apk_Urls.yml -v apkpath=test/allsafe.apk -detailed
================================================================================
[Workflow-Credit] Tasked Workflow 'APK URLs Checker' Workflow-Author=Rikunj Sindhwad
================================================================================
[START] [2023-12-04T14:24:08-05:00] Task Started TaskName=Decompile APK
[Task-Info] Task is Static TaskName=Decompile APK
[Static-Task: Decompile APK] [2023-12-04T14:24:08-05:00] Executing Task
[SUCCESS] [2023-12-04T14:24:26-05:00] Task Finished TaskName=Decompile APK
[START] [2023-12-04T14:24:27-05:00] Task Started TaskName=Extract URLS
[Task-Info] Task is Static TaskName=Extract URLS
[Static-Task: Extract URLS] [2023-12-04T14:24:27-05:00] Executing Task
[SUCCESS] [2023-12-04T14:24:32-05:00] Task Finished TaskName=Extract URLS
[Task-Info] There are required Tasks Extract URLS TaskName=Remove Duplicates
[START] [2023-12-04T14:24:33-05:00] Task Started TaskName=Remove Duplicates
[Task-Info] Task is Static TaskName=Remove Duplicates
[Static-Task: Remove Duplicates] [2023-12-04T14:24:33-05:00] Executing Task
[SUCCESS] [2023-12-04T14:24:35-05:00] Task Finished TaskName=Remove Duplicates
[Task-Info] There are required Tasks Remove Duplicates TaskName=Save Results
[START] [2023-12-04T14:24:36-05:00] Task Started TaskName=Save Results
[Task-Info] Task is Static TaskName=Save Results
[Static-Task: Save Results] [2023-12-04T14:24:36-05:00] Executing Task
Check-Result: /home/kali/Workflows/hive/Save-Results/out/result.txt
IdentifiedURLS:55
[SUCCESS] [2023-12-04T14:24:37-05:00] Task Finished TaskName=Save Results
================================================================================
[Workflow-Complete] Workflow 'APK URLs Checker' Execution Complete Time Taken=30.324584125s
================================================================================
  ```
</details>


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

To get started with TaskNinja, please refer to the [Getting Started Guide](https://github.com/RikunjSindhwad/Task-Ninja/wiki/Getting-Started) in Wiki. It provides step-by-step instructions on installation and usage. You can find all workflows in [Task-Ninja-Workflows](https://github.com/RikunjSindhwad/Task-Ninja-Workflows)

## Contributing

I welcome contributions from the community! If you'd like to contribute to TaskNinja, please follow [Contribution Guidelines](CONTRIBUTING.md).

## License

TaskNinja is open-source and released under the [GPL License 3.0](LICENSE).

## Contact

If you have any questions, feedback, or need assistance, feel free to me on [linkedin](https://www.linkedin.com/in/rikunj/) or [telegram](https://t.me/R0B077)

## Join Robensive Community
Robensive community helps you to get regular updates on news/jobs/tools/courses and knowledge sharing
Join over [telegram](https://t.me/robensive)
