package main

type IPlugin interface {
	Process(event map[string]interface{}) map[string]interface{}
}

func NewProcessor() *Processor {
	return &Processor{
		plugins: []IPlugin{},
	}
}

type Processor struct {
	plugins []IPlugin
}

func (p *Processor) Process(event map[string]interface{}) map[string]interface{} {
	for _, plugin := range p.plugins {
		event = plugin.Process(event)
	}

	return event
}

func (p *Processor) AddPlugin(plugin IPlugin) {
	p.plugins = append(p.plugins, plugin)
}
