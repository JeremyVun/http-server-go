import io.ktor.network.selector.SelectorManager
import io.ktor.network.sockets.aSocket
import io.ktor.network.sockets.Socket
import io.ktor.network.sockets.openReadChannel
import io.ktor.network.sockets.openWriteChannel
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.runBlocking
import kotlinx.coroutines.withTimeout
import io.ktor.utils.io.writeStringUtf8

fun main(args: Array<String>) {
    val port = getIntArg(args, "-p", 8080)
    val isPipeliningAllowed = getBoolArg(args, "-ka", false)
    println("Starting server on port: $port - isPipeliningAllowed: $isPipeliningAllowed")

    runBlocking {
        // Selector manager will manage the sys calls for our sockets
        // Create the socket with the manager, bind to tcp, listen
        val selectorManager = SelectorManager(Dispatchers.IO)
        val serverSocket = aSocket(selectorManager)
            .tcp()
            .bind("localhost", port)
            .also {
                println("Listening on port: $port")
            }

        while (true) {
            val socket = serverSocket.accept()
            launch() {
                runCatching {
                    handleConnection(socket, isPipeliningAllowed)
                    socket.close()
                }.onFailure { println(it) }
            }
        }
    }
}

suspend fun handleConnection(socket: Socket, isPipeliningAllowed: Boolean) {
    println("Handling new connection from ${socket.remoteAddress}")
    val receiveChannel = socket.openReadChannel()
    val sendChannel = socket.openWriteChannel()

    while (true) {
        val request = withTimeout(5000L) {
            val requestLine = parseRequestLine(receiveChannel)
            val headers = parseHeaders(receiveChannel)
            val body = headers.get("content-length")?.let { contentLength ->
                parseBody(receiveChannel, contentLength.toInt())
            }

            HttpRequest(
                requestLine.scheme,
                requestLine.method,
                requestLine.path,
                headers,
                body
            )
        }

        // head of line blocking
        val response = handleRequest(request)
        sendChannel.writeStringUtf8(response.toHttpString(isPipeliningAllowed))

        if (!isPipeliningAllowed) {
            break
        }
    }
}
