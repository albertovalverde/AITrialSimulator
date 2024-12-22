# AITrialSimulator

Un simulador de juicios que pone a prueba a los abogados recreando todos los pasos de un juicio real en un formato de "Elige tu propia aventura". El sistema presenta varias etapas, incluyendo declaraciones iniciales, testimonios de testigos y argumentos finales, de una manera dinámica e interactiva. Al incorporar la entrada del usuario, Gemini genera escenarios judiciales realistas y se adapta a las estrategias y acciones de los abogados, proporcionando una experiencia de juicio desafiante e inmersiva donde cada decisión puede influir en el resultado.

![image](https://github.com/user-attachments/assets/c519fcb9-bfc0-425b-9f99-eef6540d9388)



### Requisitos previos

*   Tener instalado Go ([https://go.dev/dl/](https://go.dev/dl/)).
*   Tener una clave de API de Gemini 2.

### Obtener una clave de API de Gemini 2

1.  Inicia Google AI Studio: [https://aistudio.google.com/](https://aistudio.google.com/)
2.  Haz clic en **Obtener clave de API**.

### Ejecutar el simulador

1.  Clona el repositorio:

    ```bash
    git clone [https://github.com/albertovalverde/AITrialSimulator.git](https://github.com/albertovalverde/AITrialSimulator.git)
    cd AITrialSimulator
    ```

2.  Configura la clave de API en la variable de entorno `API_KEY`:

    ```bash
    export API_KEY=<tu_clave_api>
    ```

    *   **Importante (macOS/Linux):** Para que la variable `API_KEY` persista durante la sesión actual de la terminal, puedes añadir la línea `export API_KEY=<tu_clave_api>` a tu archivo `~/.bashrc` o `~/.zshrc`.
    *   **Importante (Windows):** Debes usar el comando `set` para configurar la variable y solo durará la sesión actual de la terminal. Para que sea permanente, debes configurarla en las variables de entorno del sistema.

3.  Ejecuta el programa:

    ```bash
    go run .
    ```

4.  Cuando se te pregunte "¿Qué tipo de juicio te gustaría simular?", proporciona un escenario judicial.

    Por ejemplo, escribe: `Quiero simular un juicio penal por robo a mano armada.`

### Compilar para Windows (Opcional)

1.  Configura las variables de entorno:

    ```bash
    set GOOS=windows
    set GOARCH=amd64
    ```

2.  Construye el ejecutable:

    ```bash
    go build -o app-windows.exe
    ```

3.  Ejecuta el ejecutable (en Windows, desde la línea de comandos en la misma carpeta):

    ```
    set API_KEY=<tu_clave_api>
    app-windows.exe
    ```

## Características

*   Simulación interactiva de juicios con formato "Elige tu propia aventura".
*   Generación dinámica de escenarios judiciales realistas con Gemini.
*   Adaptación a las estrategias y decisiones del usuario.
*   Interfaz de usuario intuitiva (se puede mejorar).
*   Potencial para la práctica y el entrenamiento de abogados y estudiantes de derecho.

## Próximas mejoras

*   Implementación de un sistema de puntuación o evaluación del desempeño.
*   Mayor variedad de escenarios judiciales predefinidos y la posibilidad de crear escenarios personalizados.
*   Integración de pruebas y evidencias virtuales (documentos, fotos, videos).
*   Mejora de la interfaz de usuario para una experiencia más inmersiva.
*   Soporte para múltiples jugadores o roles (juez, jurado, etc.).

## Contribución

Las contribuciones son bienvenidas. Si deseas contribuir al desarrollo de este proyecto, por favor, abre un *issue* para discutir la propuesta o envía directamente un *pull request* con los cambios.

## Licencia

Este proyecto se distribuye bajo la licencia MIT. Puedes encontrar una copia de la licencia en el archivo `LICENSE` (si se incluye en el repositorio).
