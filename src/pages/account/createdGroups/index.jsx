import React, { useState, useRef, useLayoutEffect } from 'react';
import { GridContent } from '@ant-design/pro-layout';
import { Menu } from 'antd';
import { history } from 'umi';
import BaseView from './components/base';
//import BindingView from './components/binding';
import NotificationView from './components/notification';
import SecurityView from './components/security';
import styles from './style.less';
import { useRequest } from 'ahooks';
import { getGroupInfo } from '@/services/getGroupInfo';
const { Item } = Menu;


const user = history.location.search;

const Groups = () => {
  const { data, reload, loading, loadMore, loadingMore } = useRequest (
    () => {
      return getGroupInfo({
        userName: user,
      });
    },
    {
      loadMore: true,
    },
  );

  const lists = data ?. list || [];


  const menuMap = {
    base: 'Basic Information',
    member: 'Group Members',
    //binding: '账号绑定',
    post: 'Posts',
  };
  const [initConfig, setInitConfig] = useState({
    mode: 'inline',
    selectKey: 'base',
  });
  const dom = useRef();

  const resize = () => {
    requestAnimationFrame(() => {
      if (!dom.current) {
        return;
      }

      let mode = 'inline';
      const { offsetWidth } = dom.current;

      if (dom.current.offsetWidth < 641 && offsetWidth > 400) {
        mode = 'horizontal';
      }

      if (window.innerWidth < 768 && offsetWidth > 400) {
        mode = 'horizontal';
      }

      setInitConfig({ ...initConfig, mode: mode });
    });
  };

  useLayoutEffect(() => {
    if (dom.current) {
      window.addEventListener('resize', resize);
      resize();
    }

    return () => {
      window.removeEventListener('resize', resize);
    };
  }, [dom.current]);

  const getMenu = () => {
    return Object.keys(menuMap).map((item) => <Item key={item}>{menuMap[item]}</Item>);
  };

  const renderChildren = () => {
    const { selectKey } = initConfig;

    switch (selectKey) {
      case 'base':
        return <BaseView />;

      case 'member':
        return <SecurityView />;

      case 'binding':
        return <BindingView />;

      case 'post':
        return <NotificationView />;

      default:
        return null;
    }
  };

  return (
    <GridContent>
      <div
        className={styles.main}
        ref={(ref) => {
          if (ref) {
            dom.current = ref;
          }
        }}
      >
        <div className={styles.leftMenu}>
          <Menu
            mode={initConfig.mode}
            selectedKeys={[initConfig.selectKey]}
            onClick={({ key }) => {
              setInitConfig({ ...initConfig, selectKey: key });
            }}
          >
            {getMenu()}
          </Menu>
        </div>
        <div className={styles.right}>
          <div className={styles.title}>{menuMap[initConfig.selectKey]}</div>
          {renderChildren()}
        </div>
      </div>
    </GridContent>
  );
};

export default Groups;