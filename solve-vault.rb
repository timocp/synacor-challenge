#! /usr/bin/env ruby
# rubocop:disable Style/NumericPredicate, Style/Documentation

State = Struct.new(:row, :col, :weight)

class Vault
  MAP = [
    [:*, 8, :-, 1],
    [4, :*, 11, :*],
    [:+, 4, :-, 18],
    [:S, :-, 9, :*]
  ].freeze

  DIRS = [
    [:north, -1,  0],
    [:east,   0,  1],
    [:south,  1,  0],
    [:west,   0, -1]
  ].freeze

  def start_state
    State.new(3, 0, 22)
  end

  def goal?(state)
    state.row == 0 && state.col == 3 && state.weight == 30
  end

  def each_successor(state, &_block)
    this_square = MAP[state.row][state.col]
    DIRS.each do |dir, dr, dc|
      succ = State.new(state.row + dr, state.col + dc, state.weight)

      # out of bounds
      next if succ.row < 0 || succ.row > 3 || succ.col < 0 || succ.col > 3

      # can't return to start
      next if MAP[succ.row][succ.col] == :S

      # determine new weight
      case this_square
      when :+ then succ.weight += MAP[succ.row][succ.col]
      when :- then succ.weight -= MAP[succ.row][succ.col]
      when :* then succ.weight *= MAP[succ.row][succ.col]
      end

      # don't explode/shatter orb
      next if succ.weight <= 0 || succ.weight >= 32_767

      # don't enter last square unless weight is correct
      next if succ.row == 0 && succ.col == 3 && succ.weight != 30

      yield succ, dir
    end
  end
end

# breadth first search
class BFS
  def self.search(problem)
    queue = [problem.start_state]
    meta = {} # map of child state back to parent states

    while (parent_state = queue.shift)
      return construct_path(parent_state, meta) if problem.goal?(parent_state)
      problem.each_successor(parent_state) do |child_state, action|
        next if meta.key?(child_state)
        meta[child_state] = [parent_state, action]
        queue.push child_state
      end
    end
  end

  def self.construct_path(state, meta)
    action_list = []
    while (row = meta[state])
      state, action = row
      action_list.unshift action
    end
    action_list
  end
end

BFS.search(Vault.new).each { |dir| puts dir }
