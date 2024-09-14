// Captura o protocolo atual (http: ou https:)
const protocol = window.location.protocol === "https:" ? "wss://" : "ws://";

// Captura o hostname (ex.: 192.168.1.100 ou localhost)
const host = window.location.hostname;

// Captura a porta atual (ex.: 8080)
const port = window.location.port ? `:${window.location.port}` : "";

// Monta a URL do WebSocket
const wsUrl = `${protocol}${host}${port}/ws`;
console.log("Server: " + wsUrl);


// Cria uma conexão WebSocket com a URL montada
const ws = new WebSocket(wsUrl);

ws.onopen = () => {
    console.log("WebSocket connection opened");
};

ws.onmessage = (event) => {
    const data = JSON.parse(event.data);
    const messageElem = document.createElement('li');
    
    
    // Se o cliente é o remetente, alinha a mensagem à direita
    if (data.issender) {
        //messageElem.style.textAlign = "right";
        messageElem.classList.add("sender-message")
        messageElem.textContent = `[${data.username}]: ${data.message}`;
    } else {
        messageElem.textContent = `[${data.username}]: ${data.message}`;
    }

    document.getElementById('messagesList').appendChild(messageElem);
};

ws.onclose = () => {
    console.log("WebSocket connection closed");
};

ws.onerror = (error) => {
    console.error("WebSocket error:", error);
};

// Lógica de envio de mensagens
document.getElementById('sendButton').addEventListener('click', () => {
    const messageInput = document.getElementById('messageInput');
    const message = messageInput.value;

    const usernameInput = document.getElementById("usernameInput")
    const username = usernameInput.value;


    if (message && ws.readyState === WebSocket.OPEN) {
        ws.send(JSON.stringify({ username: username, message: message }));
        messageInput.value = '';
    }
});
