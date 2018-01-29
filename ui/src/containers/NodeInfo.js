import React, {Component} from 'react';
import Panel from '../components/Panel'
import FormField from '../components/FormField'

class NodeInfo extends Component {
  constructor(props) {
    super(props)
    this.state = {
      nodeInfo: null
    }
  }
  componentDidMount() {
    fetch('/api/config').then((resp) => {
      return resp.json()
    }).then((data) => {
      this.setState({
        nodeInfo: data
      })
    })
  }
  render() {
    if (!this.state.nodeInfo) {
      return <span>Loading...</span>
    }
    return <Panel title="Node Settings">
      <div className="form">
        <FormField
          label="Node ID" value={this.state.nodeInfo.nodeId}
          leftCol="col-md-3" rightCol="col-md-5" />
        <FormField label="Web Port" value={this.state.nodeInfo.webPort}
          leftCol="col-md-3" rightCol="col-md-5" />
        <FormField label="Dashboard Port" value={this.state.nodeInfo.localPort}
          leftCol="col-md-3" rightCol="col-md-5" />
        <FormField label="Shared Directory" value={this.state.nodeInfo.sharedDirectory}
          leftCol="col-md-3" rightCol="col-md-5" />
      </div>
    </Panel>
  }
}

export default NodeInfo;
