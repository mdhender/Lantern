import React, { Component } from 'react';

import {
  Form,
  Header,
  Segment,
} from 'semantic-ui-react';

export default class App extends Component {
  state = {
    command: "",
    log: [],
  }

  handleClick = () => {
    this.setState({
      command: "",
      log: [
        ...this.state.log,
        "" + (this.state.log.length + 1) + "> " + this.state.command
      ],
    });
  };
  handleMessage = e => this.setState({ ...this.state, command: e.target.value });

  updateLog = message => this.setState({ log: [...this.state.log, this.state.command + " " + this.state.log.length] });

  render() {
    const { command, log } = this.state;

    return (
      <Segment>
        <Segment>
          <Header as='h1'>MUDDLE</Header>
        </Segment>
        <Segment>
          <Form>
            <Form.TextArea rows={23} value={log.slice(-23).join("\n")} />
            <Form.Input focus fluid placeholder='Enter your commands here...' onChange={this.handleMessage} value={command} />
            <Form.Button type='submit' content='Submit' onClick={this.handleClick} />
          </Form>
        </Segment>
        <Segment>
          <p>MUDDLE</p>
        </Segment>
      </Segment>
    );
  }
}
