require 'spec_helper'

describe 'win-upgrade-rollback' do
  it_behaves_like 'an installed Agent'
  it_behaves_like 'a running Agent with no errors'
  it_behaves_like 'an upgraded Agent with the expected version'
end
