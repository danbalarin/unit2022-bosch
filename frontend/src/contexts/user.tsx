import type { ReactNode } from 'react';
import {
  useEffect,
  useMemo,
  createContext,
  useContext,
  useCallback,
  useState,
} from 'react';

import { loginUser } from '../api/auth/postLogin';
import type { IUserLoginData, IUser } from '../api/types/user';

interface IAuthContext {
  accessToken: string | null;
  user: IUser | null;
  handleUserSignIn: (data: IUserLoginData) => Promise<void>;
  handleUserSignOff: () => void;
  storeSignedUser: (data: IStoreSignedUserProps) => void;
}

interface IStoreSignedUserProps {
  user: IUser;
}

const AuthContext = createContext<IAuthContext | null>(null);

const useAuthContext = () => {
  const contextValue = useContext(AuthContext);

  if (!contextValue) {
    throw new Error(
      'useAuthContext must be used within AuthContextProvider scope'
    );
  }

  return contextValue;
};

interface IAuthContextProviderProps {
  children?: ReactNode | undefined;
}

function AuthContextProvider({ children }: IAuthContextProviderProps) {
  const [user, setUser] = useState<IUser | null>(null);
  const [accessToken, setAccessToken] = useState<string | null>(null);
  useEffect(() => {
    const storedUser = localStorage.getItem('user');

    if (storedUser) {
      setUser(JSON.parse(storedUser));
    }
  }, []);

  /* User sign in */
  const handleUserSignIn = useCallback(
    async (formData: IUserLoginData) => {
      const response = await loginUser(formData);
      setAccessToken(response.token);
      localStorage.setItem('user', JSON.stringify(response));
      setUser(response);
    },
    [setAccessToken]
  );

  /* User sign off */
  const handleUserSignOff = useCallback(() => {
    setAccessToken(null);
    setUser(null);
    localStorage.removeItem('user');
  }, [setAccessToken]);

  /* Store data about signed user */
  const storeSignedUser = useCallback(
    (data: IStoreSignedUserProps) => {
      setAccessToken(data.user.token);
      localStorage.setItem('user', JSON.stringify(data.user));
      setUser(data.user);
    },
    [setAccessToken, setUser]
  );

  const contextValue: IAuthContext = useMemo(() => {
    return {
      accessToken,
      storeSignedUser,
      user: user,
      handleUserSignIn,
      handleUserSignOff,
    };
  }, [handleUserSignIn, handleUserSignOff, user, accessToken, storeSignedUser]);

  return (
    <AuthContext.Provider value={contextValue}>{children}</AuthContext.Provider>
  );
}

export { AuthContext, useAuthContext, AuthContextProvider };
