# RosenChat
An opinionated Go framework for backend applications that encourages loose-coupling.


## How to Use
RosenChat is meant to be used as a template. Follow these steps to use it as the base of your next project.

1. Clone the repository.
2. Rename the rosenchat folder to your own project name.  
    ```mv serentiy <your-project-name>```
3. Remove RosenChat's git folder and add your own.  
    ```rm -rf .git && git init```
4. Rename all occurrences of "serentiy" in the code with your project name.  
    ```find . -type f -exec sed -i "s/rosenchat/<your-project-name>/g" {} \;```  
    ```find . -type f -exec sed -i "s/RosenChat/<Your-Project-Name>/g" {} \;```
5. Run the project.  
    ```make run```  

This should start an HTTP and a gRPC server right away. You can enable/disable these servers using the configs in the .env file. Feel free to explore all the available configs.
