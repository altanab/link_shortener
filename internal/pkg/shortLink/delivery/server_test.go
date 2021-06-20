package delivery

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"linkShortener/internal/pkg/shortLink/mocks"
	pb "linkShortener/proto"
	"testing"
)

func TestCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUc := mocks.NewMockShortLinkUsecase(mockCtrl)
	testServ := ShortLinkServer{
		mockUc,
	}
	originalLink := "originalLink"
	shortenLink := "shortenLink"

	mockUc.EXPECT().Create(originalLink).Return(shortenLink, nil).Times(1)
	sl, err := testServ.Create(
		context.Background(),
		&pb.OriginalLink{
			OriginalLink: originalLink,
		})
	require.NoError(t, err)
	assert.Equal(t, shortenLink, sl.ShortenLink)

	mockUc.EXPECT().Create(originalLink).Return("", errors.New("some error")).Times(1)
	_, err = testServ.Create(
		context.Background(),
		&pb.OriginalLink{
			OriginalLink: originalLink,
		})
	require.Error(t, err)
}

func TestGet(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUc := mocks.NewMockShortLinkUsecase(mockCtrl)
	testServ := ShortLinkServer{
		mockUc,
	}
	originalLink := "originalLink"
	shortenLink := "shortenLink"

	mockUc.EXPECT().Get(shortenLink).Return(originalLink, nil).Times(1)
	ol, err := testServ.Get(
		context.Background(),
		&pb.ShortenLink{
			ShortenLink: shortenLink,
		})
	require.NoError(t, err)
	assert.Equal(t, originalLink, ol.OriginalLink)

	mockUc.EXPECT().Get(shortenLink).Return("", errors.New("some error")).Times(1)
	_, err = testServ.Get(
		context.Background(),
		&pb.ShortenLink{
			ShortenLink: shortenLink,
		})
	require.Error(t, err)
}

