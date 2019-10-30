import React, {Component} from 'react';
import { withTranslation } from 'react-i18next';
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
    const { t } = this.props;
    if (!this.state.nodeInfo) {
      return <span>Loading...</span>
    }
    return <Panel title={t("sections.node_information.node_settings")}>
      <div className="form">
        <FormField
          label={t("sections.node_information.node_id")} 
          value={this.state.nodeInfo.nodeId}
          leftCol="col-md-3" rightCol="col-md-5" />
        <FormField label={t("sections.node_information.web_port")} 
          value={this.state.nodeInfo.webPort}
          leftCol="col-md-3" rightCol="col-md-5" />
        <FormField label={t("sections.node_information.dashboard_port")} 
          value={this.state.nodeInfo.localPort}
          leftCol="col-md-3" rightCol="col-md-5" />
        <FormField label={t("sections.node_information.shared_directory")} 
          value={this.state.nodeInfo.sharedDirectory}
          leftCol="col-md-3" rightCol="col-md-5" />
      </div>
    </Panel>
  }
}

export default withTranslation()(NodeInfo);
