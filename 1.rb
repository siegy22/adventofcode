require 'minitest/autorun'

INPUT = File.read('./1.input').strip.split("\n").map(&:to_i)

def report_repair(input, count: 2, expected_sum: 2020)
  input.permutation(count).find { |result| result.sum == expected_sum }.inject(:*)
end

class Day1Test < Minitest::Test
  def test_part1
    input = [1721, 979, 366, 299, 675, 1456]
    assert_equal 514_579, report_repair(input)
  end

  def test_part2
    input = [1721, 979, 366, 299, 675, 1456]
    assert_equal 241_861_950, report_repair(input, count: 3)
  end
end

puts "Result is: #{report_repair(INPUT)}"
puts "Result part 2 is: #{report_repair(INPUT, count: 3)}"
puts
