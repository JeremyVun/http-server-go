import io.ktor.utils.io.ByteReadChannel
import io.ktor.utils.io.readUTF8Line
import io.ktor.utils.io.readFully

// GET /path HTTP/1.1
suspend fun parseRequestLine(channel: ByteReadChannel): RequestLine =
    channel.readUTF8Line()?.split(" ")?.let { tokens ->
        RequestLine(
            scheme = tokens[2],
            method = tokens[0],
            path = tokens[1],
        )
    } ?: error("Could not parse RequestLine")

suspend fun parseHeaders(channel: ByteReadChannel): Map<String, String?> {
    val headers = mutableMapOf<String, String?>()

    var line: String? = null
    do {
        line = channel.readUTF8Line()
            ?.takeIf { it.isNotBlank() }
            ?.also {
                val tokens = it?.split(":", limit = 2)
                when (tokens?.size) {
                    1 -> headers[tokens[0]] = ""
                    2 -> headers[tokens[0]] = tokens[1]
                }
            }
    } while(line != null)

    return headers
}

suspend fun parseBody(channel: ByteReadChannel, contentLength: Int): String {
    val buffer = ByteArray(contentLength)
    channel.readFully(buffer)
    return buffer.decodeToString()
}
