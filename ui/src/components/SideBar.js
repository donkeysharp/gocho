import React, {Component} from 'react';

class SideBar extends Component {
  constructor(props) {
    super(props)
    this.state = {
      currentItem: 0
    }
  }
  itemClickHandler(e) {
    let index = parseInt(e.currentTarget.dataset.index, 10)
    if (index === this.state.currentItem) {
      return
    }
    this.setState({
      currentItem: index,
    })

    if (this.props.onMenuSelected) {
      this.props.onMenuSelected(index)
    }
  }
  render() {
    let className = 'sidebar';
    if (this.props.toggle) {
      className += ' active';
    }
    return <nav className={className}>
      <div className="sidebar-header">
        {this.props.title}
      </div>
      <ul className="list-unstyled components">
        {this.props.menu.map((item, index) => {
          let className = (index === this.state.currentItem ? 'active' : '');
          return <li className={className} key={index}>
            <a data-index={index}
              href="#/"
              onClick={this.itemClickHandler.bind(this)}>
              {item.name}
            </a>
          </li>
        })}
      </ul>
    </nav>
  }
}

export default SideBar;
