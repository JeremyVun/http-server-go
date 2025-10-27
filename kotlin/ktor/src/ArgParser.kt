fun getIntArg(args: Array<String>, flag: String, default: Int): Int =
    getStringArg(args, flag, null)?.toInt() ?: default

fun getBoolArg(args: Array<String>, flag: String, default: Boolean): Boolean =
    when (getStringArg(args, flag, null)?.lowercase()) {
        "true" -> true
        else -> false
    }

fun getStringArg(args: Array<String>, flag: String, default: String?): String? =
    args.indexOf(flag)
        .takeIf { it >= 0 }
        ?.let { flagIndex -> flagIndex + 1 }
        ?.takeIf { it < args.size }
        ?.let { valueIndex -> args[valueIndex] }
        ?: default
