import React from 'react';
import { List } from 'antd';
const passwordStrength = {
  strong: <span className="strong">strong</span>,
  medium: <span className="medium">medium</span>,
  weak: <span className="weak"weak</span>,
};

const SecurityView = () => {
  const getData = () => [
    {
      title: 'password',
      description: (
        <>
          current password strength：
          {passwordStrength.strong}
        </>
      ),
      actions: [<a key="Modify">edit</a>],
    },
    {
      title: 'phone',
      description: `35*******7`,
      actions: [<a key="Modify">edit</a>],
    },
    {
      title: 'email',
      description: `*******@ufl.edu`,
      actions: [<a key="Modify">edit</a>],
    },
     {
      title: '密保问题',
      description: '未设置密保问题，密保问题可有效保护账户安全',
      actions: [<a key="Set">设置</a>],
    },
    {
      title: 'MFA 设备',
      description: '未绑定 MFA 设备，绑定后，可以进行二次确认',
      actions: [<a key="bind">绑定</a>],
      title: 'email',
      description: `*******@ufl.edu`,
      actions: [<a key="Modify">edit</a>],
    },
    
  ];

  const data = getData();
  return (
    <>
      <List
        itemLayout="horizontal"
        dataSource={data}
        renderItem={(item) => (
          <List.Item actions={item.actions}>
            <List.Item.Meta title={item.title} description={item.description} />
          </List.Item>
        )}
      />
    </>
  );
};

export default SecurityView;
