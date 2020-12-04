require 'minitest/autorun'

INPUT = File.readlines('3.input', chomp: true)
TREE = "#".freeze

def trees_encountered(input, right: 3, down: 1)
  if input.count * right > input.first.length
    input = input.map do |row|
      row * (input.count.to_f * right.to_f / input.first.length.to_f).ceil
    end
  end

  input.to_enum.with_index.count do |line, idx|
    next if idx.zero?
    next unless idx.modulo(down).zero?

    line[(idx / down) * right] == TREE
  end
end

PART2_SLOPES = [
  { right: 1, down: 1 },
  { right: 3, down: 1 },
  { right: 5, down: 1 },
  { right: 7, down: 1 },
  { right: 1, down: 2 }
]


class Day3Test < Minitest::Test
  TEST_INPUT = [
    "..##.......",
    "#...#...#..",
    ".#....#..#.",
    "..#.#...#.#",
    ".#...##..#.",
    "..#.##.....",
    ".#.#.#....#",
    ".#........#",
    "#.##...#...",
    "#...##....#",
    ".#..#...#.#",
  ].freeze

  # def test_part1
  #   assert_equal 7, trees_encountered(TEST_INPUT)
  # end

  def test_part2
    output = PART2_SLOPES.map { |args| trees_encountered(TEST_INPUT, **args) }
    assert_equal [2, 7, 3, 4, 2], output
    assert_equal 336, output.reduce(:*)
  end
end

puts "Result is: #{trees_encountered(INPUT)}"
puts "Result part 2 is: #{PART2_SLOPES.map { |args| trees_encountered(INPUT, **args) }.reduce(:*)}"
puts
