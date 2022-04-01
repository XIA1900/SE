import React, { useCallback } from 'react';
import { LogoutOutlined, SettingOutlined, UserOutlined, CrownOutlined, HeartOutlined } from '@ant-design/icons';
import { Avatar, Menu, Spin } from 'antd';
import { history, useModel } from 'umi';
import { stringify } from 'querystring';
import HeaderDropdown from '../HeaderDropdown';
import styles from './index.less';
import { outLogin } from '@/services/ant-design-pro/api';

/**
 * logout and save the url
 */
const loginOut = async () => {
  await outLogin();
  const { query = {}, pathname } = history.location;
  const { redirect } = query; // Note: There may be security issues, please note

  if (window.location.pathname !== '/user/login' && !redirect) {
    history.replace({
      pathname: '/user/login',
      search: stringify({
        redirect: pathname,
      }),
    });
  }
};

const AvatarDropdown = ({ menu }) => {
  const { initialState, setInitialState } = useModel('@@initialState');
  const { currentUser } = initialState;

  const onMenuClick = useCallback(
    (event) => {
      const { key } = event;

      if (key === 'logout') {
        setInitialState((s) => ({ ...s, currentUser: undefined }));
        //loginOut();
        history.push('/user/login');
        return;
      }

      if(key === 'created_groups') {
        history.push({
          pathname: '/account/selectGroups/created',
          search: currentUser.name,
        });
        return;
      }

      if(key === 'joined_groups') {
        history.push({
          pathname: '/account/selectGroups/joined',
          search: currentUser.name,
        });
        return;
      }

      history.push({
        pathname: `/account/${key}`,
        search: currentUser.name,
      });
      return;
    },
    [setInitialState],
  );

  const loading = (
    <span className={`${styles.action} ${styles.account}`}>
      <Spin
        size="small"
        style={{
          marginLeft: 8,
          marginRight: 8,
        }}
      />
    </span>
  );

  if (!initialState) {
    return loading;
  }

  

  if (!currentUser || !currentUser.name) {
    return loading;
  }

  const menuHeaderDropdown = (
    <Menu className={styles.menu} selectedKeys={[]} onClick={onMenuClick}>
      {menu && (
        <Menu.Item key="center">
          <UserOutlined />
          Personal Center
        </Menu.Item>
      )}
      {menu && (
        <Menu.Item key="settings">
          <SettingOutlined />
          Settings
        </Menu.Item>
      )}
      {menu && (
        <Menu.Item key="created_groups">
          <CrownOutlined />
          Created Groups
        </Menu.Item>
      )}
      {menu && (
        <Menu.Item key="joined_groups">
          <HeartOutlined />
          Joined Groups
        </Menu.Item>
      )}
      {menu && <Menu.Divider />}

      <Menu.Item key="logout">
        <LogoutOutlined />
        Log out
      </Menu.Item>
    </Menu>
  );
  return (
    <HeaderDropdown overlay={menuHeaderDropdown}>
      <span className={`${styles.action} ${styles.account}`}>
        {/* <Avatar size="small" className={styles.avatar} src={currentUser.avatar} alt="avatar" /> */}
        <img src='http://192.168.3.132:10010/resources/userfiles/exia/exia.png'/>
        <span className={`${styles.name} anticon`}>{currentUser.name}</span>
      </span>
    </HeaderDropdown>
  );
};

export default AvatarDropdown;
