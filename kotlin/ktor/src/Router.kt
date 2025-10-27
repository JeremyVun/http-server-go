
suspend fun handleRequest(request: HttpRequest): HttpResponse {
    println("Handling request: $request")

    return HttpResponse(
        scheme = request.scheme,
        statusCode = HttpStatusCode.OK,
        headers = emptyMap(),
        body = "hello"
    )
}
