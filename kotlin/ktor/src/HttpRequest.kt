data class HttpRequest(
    val scheme: String,
    val method: String,
    val path: String,
    val headers: Map<String, String?>,
    val body: String?
)
