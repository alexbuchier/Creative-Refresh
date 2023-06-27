//1
object CreativeRefresh {
  def main(args: Array[String]): Unit = {
    println("Creative Refresh Beginning...")
    println("-----------------------------")

    //2
    val itemList: List[String] = List("Xyz", "abc", "def", "ghi")
    itemList.foreach(println)
    println("-----------------------------")

    //3
    itemList.map(item => println(s"This is modified $item"))
    println("-----------------------------")

    //4
    val itemListReverse = itemList.reverse
    itemListReverse.foreach(println)
    println("-----------------------------")

    //5
    val itemListTuple = itemList.zipWithIndex
    itemListTuple.map {
      case (item, index) => println(s"Index: $index, Item: $item")
    }
    println("-----------------------------")

    //6
    itemList.filter(item => item.contains("a") || item.contains("e")).foreach(println)
    println("-----------------------------")

    //7
    val doubleList = itemList.map(item => item * 2)
    doubleList.foreach(println)
    println("-----------------------------")

    //8
    val concatItemList = itemList.flatMap(item => item.split(""))
    concatItemList.foreach(println)
    println("-----------------------------")

    //9
    val validList = itemList.map(item => item.matches("[a-z]+"))
    validList.foreach(println)
    println("-----------------------------")

    //10
    val evenItemList = itemList.zipWithIndex.collect {
      case (item, index) if (index % 2 == 0) => (item, index)
    }
    evenItemList.foreach(println)
    println("-----------------------------")

    //11
    val sortedItemList = itemList.sortBy(item => item)
    sortedItemList.foreach(println)
    println("-----------------------------")

    //12
    val distinctItemList = itemList.distinct
    distinctItemList.foreach(println)
    println("-----------------------------")

    //13
    val concatItemListInterleave = itemList.interleave(itemList)
    concatItemListInterleave.foreach(println)
    println("-----------------------------")

    //14
    val formattedItemList = itemList.map(item => item.toUpperCase)
    formattedItemList.foreach(println)
    println("-----------------------------")

    //15
    val reversedItemList = itemList.sliding(2, 2).map { case Array(x, y) => y + x }
    reversedItemList.foreach(println)
    println("-----------------------------")

    //16
    val groupedItemList = itemList.groupBy(item => item)
    groupedItemList.foreach(println)
    println("-----------------------------")

    //17
    val takeItemList = itemList.take(2)
    takeItemList.foreach(println)
    println("-----------------------------")

    //18
    val droppedItemList = itemList.drop(2)
    droppedItemList.foreach(println)
    println("-----------------------------")

    //19
    val sumItemList = itemList.foldLeft(0)((acc, item) => acc + item.length)
    println(sumItemList)
    println("-----------------------------")

    //20
    val containsItem = itemList.exists(item => item == "ghi")
    println(containsItem)
    println("-----------------------------")

    //21
    val reducedItemList = itemList.reduce(_ + _)
    println(reducedItemList)
    println("-----------------------------")

    //22
    val mixedItemList = itemList.scanLeft("")(_ + _)
    mixedItemList.foreach(println)
    println("-----------------------------")

    //23
    val allMatchedItemList = itemList.forall(item => item.length == 3)
    println(allMatchedItemList)
    println("-----------------------------")

    //24
    val countItemList = itemList.count(item => item.length == 3)
    println(countItemList)
    println("-----------------------------")

    //25
    val maxItemIndex = itemList.zipWithIndex.maxBy(_._1.length)._2
    println(maxItemIndex)
    println("-----------------------------")

    //26
    val sumItemListPart = itemList.map(_.length()).sum
    println(sumItemListPart)
    println("-----------------------------")

    //27
    val minItemIndex = itemList.zipWithIndex.minBy(_._1.length)._2
    println(minItemIndex)
    println("-----------------------------")

    //28
    val itemListReversePart = itemList.zipWithIndex.map {
      case (item, index) if index % 2 == 0 => item
    }
    itemListReversePart.foreach(println)
    println("-----------------------------")

    //29
    val itemListPart = itemList.zipWithIndex.filter {
      case (item, index) => index % 2 != 0
    }
    itemListPart.foreach(println)
    println("-----------------------------")

    //30
    val itemListIntermittent = itemList.zipWithIndex.flatMap {
      case (item, index) if index % 2 == 0 => Some(item)
      case _ => None
    }
    itemListIntermittent.foreach(println)
    println("-----------------------------")

    println("Creative Refresh completed!")
  }
}