# Application Monitor plugin

#### Plugin arguments

* msrCpu boolean: If true, the plugin will collect CPU metrics
* msrVirMem boolean: If true, the plugin will collect Virtual Memory metrics
* appName string: The name of the application which is to be monitored

##### Configuration:

```
[[inputs.application_monitor]]
    ## Whether to include CPU measurements from the file
    msrCPU = true
    ## Whether to include IO measurements from the file
    msrIO = true
    ## The name of the application you want to monitor
    appName = "Monitor"
```

#### Description
The Application Monitor plugin monitors software on the Application layer of the OSI model.
It is designed to collect performance metrics from any application to determine its performance.
It is specifically designed to work with the FORCE framework, but can be altered to work with any
other application.