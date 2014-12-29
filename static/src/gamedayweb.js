
var data = { Gameday: "gid_2014_06_04_seamlb_atlmlb_1" };

var GameBox = React.createClass({
    render: function() {
        return (
            <div className="gameBox">
                This is where the boxscore will go.
                <GameForm/>
                <GameBoxscore data={this.props.data}></GameBoxscore>
            </div>
        );
    }
});

var GameForm = React.createClass({
    render: function() {
        return (
            <div className="gameForm">
                This is where users will select a date and team.
            </div>
        );
    }
});

var GameBoxscore = React.createClass({
    render: function() {
        return (
            <div className="gameBoxscore">
                <h3>{this.props.data.Gameday}</h3>
                {this.props.children}
            </div>
            );
    }
});

React.render(
    <GameBox data={data}/>,
    document.getElementById('content')
);
