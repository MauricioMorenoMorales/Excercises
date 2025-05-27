class ListNode {
  val: number
  next: ListNode | null
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val
    this.next = next === undefined ? null : next
  }
}

function findTheDifference(s: string, t: string): string {
  
  for (let i = 0; i < s.length || i < t.length; i++) {
    if (s[i] !== t[i]) return t[i]
  }
  return ""
}