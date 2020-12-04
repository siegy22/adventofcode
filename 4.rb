require 'minitest/autorun'
require 'ostruct'

INPUT = File.readlines('4.input', "\n\n", chomp: true)

class Passport < OpenStruct
  REQUIRED_FIELDS = %w(pid ecl hcl hgt eyr iyr byr).freeze
  VALID_EYE_COLORS = %w(amb blu brn gry grn hzl oth)

  def self.parse(str)
    str = str.gsub("\n", " ")
    kwargs = str.split(" ").each_with_object({}) do |obj, memo|
      key, value = obj.split(":")
      memo[key] = value
    end
    new(kwargs)
  end

  def valid_part2?
    required_fields_set? && fields_valid?
  end

  def required_fields_set?
    REQUIRED_FIELDS.all? { |field| self[field] }
  end
  alias_method :valid_part1?, :required_fields_set?

  def fields_valid?
    to_h.keys.all? do |field|
      send("valid_#{field}?")
    end
  end

  def valid_byr?
    (1920..2002).include?(byr.to_i)
  end

  def valid_iyr?
    (2010..2020).include?(iyr.to_i)
  end

  def valid_eyr?
    (2020..2030).include?(eyr.to_i)
  end

  def valid_hgt?
    return unless /\A\d+(cm|in)\z/.match(hgt)

    return (59..76).include?(hgt.to_i) if hgt.include?("in")
    return (150..193).include?(hgt.to_i) if hgt.include?("cm")
  end

  def valid_hcl?
    /\A\#\h{6}\z/.match(hcl)
  end

  def valid_ecl?
    VALID_EYE_COLORS.include?(ecl)
  end

  def valid_pid?
    /\A\d{9}\z/.match(pid)
  end

  def valid_cid?
    true
  end
end

def number_of_valid_passports(passports, meth:)
  passports.count do |passport|
    Passport.parse(passport).public_send(meth)
  end
end

class Day4Test < Minitest::Test
  TEST_INPUT = [
    "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm",
    "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\nhcl:#cfa07d byr:1929",
    "hcl:#ae17e1 iyr:2013\neyr:2024\necl:brn pid:760753108 byr:1931\nhgt:179cm",
    "hcl:#cfa07d eyr:2025 pid:166559648\niyr:2011 ecl:brn hgt:59in"
  ].freeze

  TEST_INPUT_PART2 = [
    "eyr:1972 cid:100\\nhcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
    "iyr:2019\nhcl:#602927 eyr:1967 hgt:170cm\necl:grn pid:012533040 byr:1946",
    "hcl:dab227 iyr:2012\necl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
    "hgt:59cm ecl:zzz\neyr:2038 hcl:74454a iyr:2023\npid:3556412378 byr:2007",
    "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\nhcl:#623a2f",
    "eyr:2029 ecl:blu cid:129 byr:1989\niyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
    "hcl:#888785\nhgt:164cm byr:2001 iyr:2015 cid:88\npid:545766238 ecl:hzl\neyr:2022",
    "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"
  ].freeze

  def test_part1
    assert_equal 2, number_of_valid_passports(TEST_INPUT, meth: :valid_part1?)
  end

  def test_part2
    assert_equal 4, number_of_valid_passports(TEST_INPUT_PART2, meth: :valid_part2?)
  end
end

puts "Result is: #{number_of_valid_passports(INPUT, meth: :valid_part1?)}"
puts "Result part 2 is: #{number_of_valid_passports(INPUT, meth: :valid_part2?)}"
puts
