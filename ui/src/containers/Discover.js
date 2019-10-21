import React, {Component} from 'react';
import { withTranslation } from 'react-i18next';
import Panel from '../components/Panel';
import NodeList from '../components/NodeList';
import NodeDetails from '../components/NodeDetails';


class Discover extends Component {
  constructor(props) {
    super(props)
    this.state = {
      nodes: [],
      currentNode: -1,
    }
  }
  retrieveData() {
    fetch('/api/nodes').then((resp) => {
      return resp.json()
    }).then((data) => {
      this.setState({
        nodes: data
      })
    })
  }
  componentDidMount() {
    this.retrieveData()
  }
  nodeSelectedHandler(index) {
    this.setState({
      currentNode: index,
    });
  }
  render() {
    const { t } = this.props;
    let detailsBody = <span>{t('sections.discover.no_node_selected')}</span>
    if (this.state.currentNode !== -1) {
      detailsBody = <NodeDetails
        node={this.state.nodes[this.state.currentNode]}
      />
    }
    return <Panel title={t("sections.discover.auto_discovery")}>
      <div className="row">
        <div className="col-md-3">
          <NodeList
            nodes={this.state.nodes}
            onNodeSelected={this.nodeSelectedHandler.bind(this)} />
        </div>
        <div className="col-md-9">
          {detailsBody}
        </div>
      </div>
    </Panel>
  }
}

export default withTranslation()(Discover);
