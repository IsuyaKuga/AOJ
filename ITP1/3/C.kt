fun main(argv: Array<String>) {
  val list: MutableList<String> = mutableListOf()

  do {
    val inputs = readLine()!!.split(' ').map{x->x.toInt()}.sorted()
    list.add("${inputs[0]} ${inputs[1]}")
  } while (!(inputs[0] == 0 && inputs[1] == 0))

  for (i in list.indices.toList().dropLast(1)) {
    println("${list[i]}")
  }
}
