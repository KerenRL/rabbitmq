class WebSocketService {
    private socket: WebSocket | null = null;
    private listeners: ((message: string) => void)[] = [];
  
    connect(url: string) {
      this.socket = new WebSocket(url);
  
      this.socket.onmessage = (event) => {
        const message = event.data;
        this.listeners.forEach((listener) => listener(message));
      };
  
      this.socket.onclose = () => {
        console.log("WebSocket cerrado. Reconectando...");
        setTimeout(() => this.connect(url), 5000); // Reconectar automáticamente
      };
    }
  
    addListener(listener: (message: string) => void) {
      this.listeners.push(listener);
    }
  
    removeListener(listener: (message: string) => void) {
      this.listeners = this.listeners.filter((l) => l !== listener);
    }
  }
  
  const websocketService = new WebSocketService();
  export default websocketService;