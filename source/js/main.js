var React = require("react");
var _ = require('lodash');
var oboe = require("oboe");

var User = React.createClass({
  render() {
    return(
        <li>{this.props.data.Name}</li>
        );
  }
});

var UserList = React.createClass({
  load() {
    oboe(this.props.url)
      .done( data => this.setState({users: data}) )
      .fail( (error) => console.error(error) )
  },
  getInitialState() {
    return {
      users: []
    };
  },
  componentDidMount() {
    this.load();
  },
  render() {
    var userlist = this.state.users.map( (user, index) => <User data={user} key={index} />)
    return(
        <ul>
          {userlist}
        </ul>
      );
  }
});

window.renderUserList = function(id) {
  React.render(
      <UserList url={`/user/list?room=${id}`} />,
      document.getElementById('userlist')
      );
}

