defstruct [:val, :next]
  defstruct [:val, :next]
end

defmodule LinkedListIntersection do
  def get_intersection_node_mapset head_a, head_b do
    nodes_a = collect_nodes(head_a, MapSet.new())

    find_intersection head_b, nodes_a
  end

  defp collect_nodes nil, acc do acc end
  defp collect_nodes node, acc  do
    collect_nodes(node.next, MapSet.put(acc, node))
  end

  defp find_intersection(nil, _nodes_set), do: nil
  defp find_intersection(node, nodes_set) do
    if MapSet.member?(nodes_set, node) do
      node
    else
      find_intersection node.next, nodes_set
    end
  end
end
