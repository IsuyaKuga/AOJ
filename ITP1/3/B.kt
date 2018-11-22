fun main(argv: Array<String>) {
  val list: MutableList<Int> = mutableListOf()

  do {
    val n = readLine()!!.toInt()
    list.add(n)
  } while (n != 0)

  for (i in list.indices.toList().dropLast(1)) {
    println("Case ${i+1}: ${list[i]}")
  }

}
