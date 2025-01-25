import { Email } from '../../models/email'
import { User, UserProps, UserStatuses, UserTypes } from '../user/user'

const BaseProps = (): UserProps => ({
  name: 'John Doe',
  email: new Email('john@mail.com'),
  createdAt: new Date(),
  modifiedAt: new Date(),
  userType: UserTypes.Consultant,
  password: {
    hash: 'foo',
    salt: 'bar',
  },
})

export const UserDummy = (partial: Partial<UserProps> = {}): User => {
  return new User({
    ...BaseProps(),
    status: UserStatuses.Active,
    ...partial,
  })
}

export const DeletedUserDummy = (partial: Partial<UserProps> = {}): User => {
  return new User({
    ...BaseProps(),
    status: UserStatuses.Deleted,
    deletedAt: new Date(),
    ...partial,
  })
}
