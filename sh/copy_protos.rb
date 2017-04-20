#!/usr/bin/env ruby

require "fileutils"

class Copier
  ROOT_DIR = File.expand_path(File.join(File.dirname(__FILE__), ".."))

  def run
    files      = Dir.glob(from_path_glob)
    code_files = files.select{|x| is_code?(x) }
    code_files.each do |f| copy_file(f) end
  end

  private

  def copy_file(file)
    # puts "copy from #{file} to #{dest_path(file)}..."
    FileUtils.mkdir_p(File.dirname(dest_path(file)))
    FileUtils.cp(file, dest_path(file))
  end

  def dest_path(file)
    file.gsub(from_path_folder, to_path_folder)
  end

  def to_path_folder
    File.join(
      ROOT_DIR,
      "src",
      "proto_out"
    )
  end

  def from_path_folder
    File.join(
      ROOT_DIR,
      "src",
      "plz-out/gen/proto"
    )
  end

  def from_path_glob
    File.join(
      from_path_folder,
      "**/**/**"
    )
  end

  def is_code?(file)
    non_code_exts.each do |ext|
      if File.extname(file) == ext
        return false
      end

      if File.directory?(file)
        return false
      end
    end
    return true
  end

  def non_code_exts
    [".a", ".proto"]
  end
end

Copier.new.run