require 'minitest/autorun'

INPUT = File.readlines('2.input', chomp: true)

class Password
  def self.parse(str)
    frequency, character, password = str.split(" ")
    new(password, character.tr(":", ""), Range.new(*frequency.split("-").map(&:to_i)))
  end

  def initialize(password, character, frequency)
    @password = password
    @character = character
    @frequency = frequency
  end

  def valid?
    @frequency.include?(@password.count(@character))
  end
end

# The shopkeeper suddenly realizes that he just accidentally explained the password policy rules
# from his old job at the sled rental place down the street! The Official Toboggan Corporate Policy
# actually works a little differently.
class Part2Password
  def self.parse(str)
    positions, character, password = str.split(" ")
    new(password, character.tr(":", ""), positions.split("-").map(&:to_i))
  end

  def initialize(password, character, positions)
    @password = password
    @character = character
    @positions = positions
  end

  def valid?
    @positions.count { |pos| @password[pos - 1] == @character } == 1
  end
end

def number_of_valid_passwords(passwords, klass: Password)
  passwords.count do |pw|
    klass.parse(pw).valid?
  end
end

class Day2Test < Minitest::Test
  TEST_INPUT = [
    "1-3 a: abcde",
    "1-3 b: cdefg",
    "2-9 c: ccccccccc"
  ].freeze

  def test_part1
    assert_equal 2, number_of_valid_passwords(TEST_INPUT)
    assert Password.parse(TEST_INPUT[0]).valid?
    refute Password.parse(TEST_INPUT[1]).valid?
    assert Password.parse(TEST_INPUT[2]).valid?
  end

  def test_part2
    assert_equal 1, number_of_valid_passwords(TEST_INPUT, klass: Part2Password)
    assert Part2Password.parse(TEST_INPUT[0]).valid?
    refute Part2Password.parse(TEST_INPUT[1]).valid?
    refute Part2Password.parse(TEST_INPUT[2]).valid?
  end
end

puts "Result is: #{number_of_valid_passwords(INPUT)}"
puts "Result part 2 is: #{number_of_valid_passwords(INPUT, klass: Part2Password)}"
puts
