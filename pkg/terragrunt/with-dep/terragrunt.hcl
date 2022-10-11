terraform {
  source = "../src"
}

dependencies {
  paths = ["../../vpc"]
}

dependency "sg" {
  config_path = "../sg"

  mock_outputs = {
    id = "fooBar"
  }
}

inputs = {
  sgID = dependency.sg.outputs.id
}
