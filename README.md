# AITrialSimulator

A trial simulator that tests lawyers by recreating every step of a real trial in a "Choose Your Own Adventure" format. The system features various stages, including opening statements, witness testimonies, and closing arguments, in a dynamic and interactive way. By incorporating user input, Gemini generates realistic court scenarios and adapts to lawyers' strategies and actions, providing a challenging and immersive trial experience where every decision can influence the outcome.


![image](https://github.com/user-attachments/assets/c519fcb9-bfc0-425b-9f99-eef6540d9388)

### Prerequisites

*   Go installed ([https://go.dev/dl/](https://go.dev/dl/)).
*   A Gemini 2 API key.

### Obtaining a Gemini 2 API Key

1.  Open Google AI Studio: [https://aistudio.google.com/](https://aistudio.google.com/)
2.  Click **Get API Key**.

### Running the Simulator

1.  Clone the repository:

    ```bash
    git clone [https://github.com/albertovalverde/AITrialSimulator.git](https://github.com/albertovalverde/AITrialSimulator.git)
    cd AITrialSimulator
    ```

2.  Set the API key in the `API_KEY` environment variable:

    ```bash
    export API_KEY=<your_api_key>
    ```

    *   **Important (macOS/Linux):** To persist the `API_KEY` variable during the current terminal session, you can add the line `export API_KEY=<your_api_key>` to your `~/.bashrc` or `~/.zshrc` file.
    *   **Important (Windows):** Use the `set` command to set the variable, which will only last for the current terminal session. To make it permanent, configure it in the system's environment variables.

3.  Run the program:

    ```bash
    go run .
    ```

4.  When prompted, "What type of trial would you like to simulate?", provide a court scenario.

    For example, type: `I want to simulate a criminal trial for armed robbery.`

### Compiling for Windows (Optional)

1.  Set the environment variables:

    ```bash
    set GOOS=windows
    set GOARCH=amd64
    ```

2.  Build the executable:

    ```bash
    go build -o app-windows.exe
    ```

3.  Run the executable (on Windows, from the command line in the same folder):

    ```
    set API_KEY=<your_api_key>
    app-windows.exe
    ```

## Features

*   Interactive trial simulation in a "Choose Your Own Adventure" format.
*   Dynamic generation of realistic court scenarios with Gemini.
*   Adapts to user strategies and decisions.
*   Intuitive user interface (can be improved).
*   Potential for practice and training for lawyers and law students.

## Upcoming developments

*   Implementation of a scoring or performance evaluation system.
*   Greater variety of predefined court scenarios and the ability to create custom scenarios.
*   Integration of virtual evidence (documents, photos, videos).
*   Improved user interface for a more immersive experience.
*   Support for multiple players or roles (judge, jury, etc.).

## Contribution

Contributions are welcome. If you wish to contribute to the development of this project, please open an *issue* to discuss the proposal or directly submit a *pull request* with your changes.



