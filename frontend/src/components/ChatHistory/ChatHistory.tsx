import React, {Component} from 'react'

interface ChatHistoryProps {
  chatHistory: { data: string }[];
}

class ChatHistory extends Component<ChatHistoryProps> {
  render(): React.ReactNode {
    const message = this.props.chatHistory.map((msg, index) => (
      <p key={index}>{msg.data}</p>
    ));

    return (
      <div className="ChatHistory">
        <h2>Chat History</h2>
        {message}
      </div>
    );
  }
}

export default ChatHistory;