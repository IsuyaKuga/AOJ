fun main(args: Array<String>) {
  val x = readLine()!!.toInt()
  val h = x / 3600
  val m = x % 3600 / 60
  val s = x % 3600 % 60
  println("${h}:${m}:${s}")
}
