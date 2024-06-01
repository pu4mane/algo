// package linkedlist

// type Node struct {
// 	val  int
// 	next *Node
// }

// type LinkedList struct {
// 	count int
// 	head  *Node
// 	tail  *Node
// }

// func (l *LinkedList) Add(val int) {
// 	newNode := &Node{
// 		val:  val,
// 		next: nil,
// 	}

// 	if l.head == nil {
// 		l.head = newNode
// 		l.tail = newNode
// 	} else {
// 		l.tail.next = newNode
// 		l.tail = newNode
// 	}

// 	l.count++
// }

// func (l *LinkedList) Remove(element int) (int, bool) {
// 	current := l.head
// 	var previous *Node

// 	index := 0
// 	for current != nil {
// 		// find element
// 		if current.val == element {

// 			// is not first node
// 			if previous != nil {

// 				// before: Head -> 1 -> 2 -> 3 --> null
// 				// after:  Head -> 1 -> 2 -------> null
// 				previous.next = current.next

// 				// set last node if current already is last
// 				if current.next == nil {
// 					l.tail = previous
// 				}
// 			} else {
// 				// before: Head -> 3 -> 5
// 				// after:  Head ------> 5

// 				// Head -> 3 -> null
// 				// Head ------> null

// 				l.head = current.next
// 				// check is empty list
// 				if l.head == nil {
// 					l.tail = nil
// 				}
// 			}

// 			l.count--
// 			return index, true
// 		}

// 		index++
// 		previous = current
// 		current = current.next
// 	}

// 	return index, false
// }

// func (l *LinkedList) Contains(element int) (int, bool) {
// 	current := l.head

// 	index := 0
// 	for current != nil {
// 		if current.val == element {
// 			return index, true
// 		}

// 		current = current.next
// 		index++
// 	}

//		return index, false
//	}
package linkedlist

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Read(index int) (int, bool) {
	// Начинаем с первого узла списка
	currentNode := l.head
	currentIndex := 0

	for currentIndex < index {
		// Последовательно переходим по ссылкам,
		// пока не доберемся до индекса искомого узла
		currentNode = currentNode.next
		currentIndex++
	}

	// Если мы вышли за пределы списка, значит,
	// нужного индекса в нем нет, поэтому возвращаем пil
	if currentNode == nil {
		return 0, false
	}

	return currentNode.data, true
}

func (l *LinkedList) IndexOf(data int) (int, bool) {
	// Начинаем с первого узла списка
	currentNode := l.head
	currentIndex := 0

	// Обнаружив искомые данные, возвращаем их
	for currentNode != nil {
		if currentNode.data == data {
			return currentIndex, true
		}

		// В противном случае переходим к следующему узлу
		currentNode = currentNode.next
		currentIndex++
	}

	// Если мы проходим весь список,
	// не обнаружив искомых данных, возвращаем пil
	return 0, false
}

func (l *LinkedList) InsertAtIndex(index, data int) {
	// Создаем новый узел с указанным значением
	newNode := &Node{
		data: data,
		next: nil,
	}
	// Если мы вставляем значение в начало списка
	if index == 0 {
		// Делаем так, чтобы наш новый узел ссылался на тот,
		// который был первым ранее
		newNode.next = l.head
		// Устанавливаем новый узел в качестве первого узла списка
		l.head = newNode
		return
	}

	// Если мы вставляем значение не в начало списка

	currentNode := l.head
	currentIndex := 0

	// Сначала обращаемся к узлу, который находится
	// перед местом вставки нового узла
	for currentIndex < index-1 {
		currentNode = currentNode.next
		currentIndex++
	}

	// Делаем так, чтобы ссылка нового узла указывала на следующий узел
	newNode.next = currentNode.next
	// Изменяем ссылку предыдущего узла,
	// чтобы она указывала на новый узел
	currentNode.next = newNode
}

func (l *LinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		// Просто устанавливаем в качестве первого узла тот,
		// который сейчас второй
		l.head = l.head.next
		return
	}

	currentNode := l.head
	currentIndex := 0

	// Сначала находим узел, предшествующий тому,
	// который мы хотим удалить, и называем его текущим currentNode
	for currentIndex < index-1 {
		currentNode = currentNode.next
		currentIndex++
	}

	// Находим узел, следующий за удаляемым
	пodeAfterDeletedNode := currentNode.next.next
	// Меняем ссылку current_node, чтобы она указывала на узел,
	// следующий за удаляемым, исключая удаляемый узел из списка
	currentNode.next = пodeAfterDeletedNode.next
}
