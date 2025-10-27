import kotlin.text.StringBuilder

data class HttpResponse(
    val scheme: String,
    val statusCode: HttpStatusCode,
    val headers: Map<String, String?>,
    val body: String?
) {
    fun toHttpString(isPipeliningAllowed: Boolean): String = StringBuilder().apply {
        append("$scheme ${statusCode.code} ${statusCode.reasonPhrase}\r\n")
        headers.toMutableMap()
            .apply {
                put("content-length", body?.length?.toString() ?: "0")
                put("connection", if (isPipeliningAllowed) { "keep-alive" } else { "close" })
            }.forEach { (key, value) ->
                append("key: value\r\n")
            }

        append("\r\n")
        body?.let { append(body) }
    }.toString()
}

enum class HttpStatusCode(val code: Int, val reasonPhrase: String) {
    OK(200, "OK"),
    NOT_FOUND(404, "Not Found"),
    INTERNAL_SERVER_ERROR(500, "Internal Server Error")
}
